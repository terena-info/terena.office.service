package validations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gerrors "github.com/terena-info/terena.godriver/gerror"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateParamObjectId(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param(key)
		if !primitive.IsValidObjectID(ID) {
			gerrors.Panic(http.StatusBadRequest, gerrors.E{Message: "validate_object_id", ErrorCode: "4000"})
		}
		ctx.Next()
	}
}
