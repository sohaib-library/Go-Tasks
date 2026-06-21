package middleware

import (
	"log"
	"net/http"
	"strings"
	"web-server/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
		log.Print("Authorization header is missing")
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
		log.Print("Invalid authorization format")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := utils.VerifyJWT(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		log.Print("Invalid Token")
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok || userIDFloat <= 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		log.Print("Invalid user_id claim")
		return
	}

	ctx.Set("user_id", int(userIDFloat))
	ctx.Next()
}
