package repositories

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	gerrors "github.com/terena-info/terena.godriver/gerror"
	"github.com/terena-info/terena.godriver/gomgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"terena.office/src/configs"
	"terena.office/src/models"
)

type AdminRepositories interface {
	FindById(primitive.ObjectID) models.Admin
	FetchAllAdmins(*gin.Context, interface{})
	Login(string, string) string
	_GenerateAccessToken(primitive.ObjectID) string
	VerifyAccessToken(string) (models.Admin, error)
	GetAuth(*gin.Context) models.Admin
}

type _Adaptor struct {
	adminOrm gomgo.OrmInterface
}

func (adapter _Adaptor) GetAuth(ctx *gin.Context) models.Admin {
	authData, _ := ctx.Get("auth")
	adminData := authData.(models.Admin)
	return adminData
}

func (adaptor *_Adaptor) FindById(id primitive.ObjectID) models.Admin {
	var admin []models.Admin
	adaptor.adminOrm.FindById(id).Unset("password").Decode(&admin).ErrorMessage(configs.ADMIN_NOT_FOUND)
	return admin[0]
}

func (adaptor *_Adaptor) FetchAllAdmins(ctx *gin.Context, output interface{}) {

}

type AdminJWTPayload struct {
	UserId primitive.ObjectID
	jwt.RegisteredClaims
}

func (adaptor *_Adaptor) Login(email, password string) string {
	var admin []models.Admin
	adaptor.adminOrm.Match(bson.M{"email": email}).Decode(&admin).ErrorMessage(configs.ADMIN_NOT_FOUND)

	if err := bcrypt.CompareHashAndPassword([]byte(admin[0].Password), []byte(password)); err != nil {
		gerrors.Panic(http.StatusBadRequest, gerrors.E{Message: "email_or_password_incorrect"})
	}

	token := adaptor._GenerateAccessToken(admin[0].ID)
	return token
}

func (adaptor *_Adaptor) _GenerateAccessToken(userId primitive.ObjectID) string {
	var payload AdminJWTPayload
	payload.UserId = userId

	payload.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 100000)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	key := []byte(configs.Env.JWT_SECRET_KEY) // Secret key
	tokenString, err := token.SignedString(key)

	if err != nil {
		gerrors.Panic(http.StatusBadRequest, gerrors.E{Message: err.Error()})
	}

	return tokenString
}

// Verify access token
func (adaptor *_Adaptor) VerifyAccessToken(token string) (models.Admin, error) {
	payload := &AdminJWTPayload{}
	var admin []models.Admin

	tokenString, err := jwt.ParseWithClaims(token, payload, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.Env.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return admin[0], errors.New("invalid_token")
	}

	if !tokenString.Valid {
		return admin[0], errors.New("invalid_signature")
	}

	adminData := adaptor.FindById(payload.UserId)
	if adminData.Email == "" {
		return adminData, errors.New("user_not_found")
	}

	return adminData, nil
}

func AdminRepository() AdminRepositories {
	var adaptor AdminRepositories = &_Adaptor{
		adminOrm: gomgo.New(context.TODO(), models.AdminModelName),
	}
	return adaptor
}
