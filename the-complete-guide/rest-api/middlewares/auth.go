package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.GetHeader("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Next()
}
