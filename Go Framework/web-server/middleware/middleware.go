package middleware

import (
	"log"
	"net/http"
	"strings"
	"web-server/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	// Extract the token from the Authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email and Password"})
		log.Print("Authorization header is missing")
		return
	}

	// Split the header to get the token part
	tokenString := strings.Split(authHeader, "Bearer ")[1]

	_, err := utils.VerifyJWT(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email and Password"})
		log.Print("Invalid Token")
		return
	}

	// Token is valid, proceed with the request
	ctx.Next()
}
