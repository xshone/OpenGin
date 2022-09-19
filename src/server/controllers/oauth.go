package controllers

import (
	"net/http"
	models "opengin/server/models"
	schemas "opengin/server/schemas"
	"opengin/server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Token(ctx *gin.Context) {
	var params schemas.OAuthLogin
	var user models.User

	if err := ctx.ShouldBind(&params); err == nil {
		models.GetDB().Where("username=?", strings.ToLower(params.Username)).First(&user)
	}

	if user.ID != 0 && user.PasswordHash == utils.HashPassword(params.Password, user.PasswordSalt) {
		token, err := utils.CreateToken(user.ID)

		if err == nil {
			ctx.JSON(http.StatusOK, schemas.OAuthTokenResponse{
				AccessToken: token,
				TokenType:   "bearer",
			})
			return
		}
	}

	http.Error(ctx.Writer, "Failed to login", http.StatusUnauthorized)
}
