package handler

import (
	"log"
	"net/http"

	"web-server/models"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      User login
// @Description  Authenticate a user and return a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.Login  true  "Login credentials"
// @Success      200          {object}  map[string]string  "token"
// @Failure      400          {object}  map[string]string  "Invalid login request"
// @Failure      401          {object}  map[string]string  "Invalid email or password"
// @Router       /login [post]
func (h *Handler) Login(ctx *gin.Context) {

	var login models.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
		log.Print(err.Error())
		return
	}

	token, err := h.Authuser.Login(login)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"token":   token,
	})
}
