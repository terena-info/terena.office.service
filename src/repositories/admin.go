package repositories

import (
	"errors"
	"net/http"
	"time"

	"github.com/bankonly/goginhandlers/src/configs"
	"github.com/bankonly/goginhandlers/src/models"
	"github.com/bankonly/goutils/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AdminInterface interface {
	Login(string, string) string // (Email, Password)
	SeedAdmin() string           // (Email, Password)
	GenerateAccessToken(primitive.ObjectID) string
	VerifyAccessToken(string) (models.Admin, error)
	FindByID(primitive.ObjectID) models.Admin
	FindOne(interface{}) (models.Admin, bool)
	GetAuth(*gin.Context) models.Admin
}

type _AdminRepository struct{}

// Get auth user from token
func (repo _AdminRepository) GetAuth(c *gin.Context) models.Admin {
	auth, ok := c.Get("auth")
	if !ok {
		utils.Panic(http.StatusInternalServerError, utils.POption{Message: "No auth found"})
	}
	adminData := auth.(models.Admin)
	return adminData
}

func (repo _AdminRepository) FindOne(condition interface{}) (models.Admin, bool) {
	var admin models.Admin
	result := models.ADMIN_INSTANCE.FindOne(configs.DBContext(), condition)
	err := result.Decode(&admin)
	return admin, err == nil
}

func (repo _AdminRepository) FindByID(ID primitive.ObjectID) models.Admin {
	var admin models.Admin
	result := models.ADMIN_INSTANCE.FindOne(configs.DBContext(), bson.M{"_id": ID})
	result.Decode(&admin)
	return admin
}

func (repo _AdminRepository) Login(email, password string) string {
	admin, exist := repo.FindOne(bson.M{"email": email})
	if !exist {
		utils.Panic(http.StatusBadRequest, utils.POption{Message: "Email or password is invalid", ErrorCode: "4000"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		utils.Panic(http.StatusBadRequest, utils.POption{Message: "Email or password is invalid", ErrorCode: "4001"})
	}

	token := repo.GenerateAccessToken(admin.ID)
	return token
}

// Verify access token
func (repo _AdminRepository) VerifyAccessToken(token string) (models.Admin, error) {
	payload := &models.AdminJWTPayload{}
	var admin models.Admin

	tokenString, err := jwt.ParseWithClaims(token, payload, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.Env.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return admin, errors.New("invalid token")
	}

	if !tokenString.Valid {
		return admin, errors.New("invalid signature")
	}

	adminData := repo.FindByID(payload.UserId)
	if adminData.Email == "" {
		return admin, errors.New("user not found")
	}

	return adminData, nil
}

func (repo _AdminRepository) GenerateAccessToken(userId primitive.ObjectID) string {
	var payload models.AdminJWTPayload
	payload.UserId = userId
	payload.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 199999)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(configs.Env.JWT_SECRET_KEY))

	if err != nil {
		utils.Panic(http.StatusInternalServerError, utils.POption{Message: "Failed to generate JWT token"})
	}
	return tokenString
}

func (repo _AdminRepository) SeedAdmin() string {
	var admin models.Admin
	models.BindCreate(&admin)

	admin.FullName = "Souksavanh Xayxomphou"
	admin.Email = "admin@gmail.com"
	pass, err := bcrypt.GenerateFromPassword([]byte("123123123"), 10)
	if err != nil {
		utils.Panic(http.StatusInternalServerError, utils.POption{Message: "Generate password failed"})
	}
	admin.Password = string(pass)

	_, err = models.ADMIN_INSTANCE.InsertOne(configs.DBContext(), admin)
	if err != nil {
		utils.Panic(http.StatusInternalServerError, utils.POption{Message: "Insert Failed"})
	}
	return "Admin Seeded"
}

func RegisterAdminRepository() AdminInterface {
	var repository AdminInterface = _AdminRepository{}
	return repository
}
