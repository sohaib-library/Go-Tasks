package handler

import (
	"log"
	"net/http"
	"regexp"
	"web-server/models"
	"web-server/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUP(ctx *gin.Context) {

	var users models.Users

	var emailRegex = regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signup request"})
		log.Print(err.Error())
		return
	}

	hashedPassword, _ := utils.EncryptPassword(users.PASSWORD)

	users.PASSWORD = hashedPassword

	if !emailRegex.MatchString(users.EMAIL) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter Valid Email",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": " SignUp Successfully"})
}
