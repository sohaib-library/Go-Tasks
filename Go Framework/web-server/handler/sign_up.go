package handler

import (
	"log"
	"net/http"
	"regexp"
	"web-server/models"
	"web-server/utils"

	"github.com/gin-gonic/gin"
)

// SignUP godoc
// @Summary      User registration
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.Users  true  "User registration details"
// @Success      200   {object}  map[string]string  "Success message"
// @Failure      400   {object}  map[string]string  "Invalid signup request or email"
// @Router       /signup [post]
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
