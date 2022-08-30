package controllers

import (
	"net/http"
	models "opengin/server/models"
	schemas "opengin/server/schemas"
	"opengin/server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Title OAuth登录
// @Tag OAuth
// @Description OAuth login
// @Param request form schemas.OAuthLogin true "_"
// @Route /v1/oauth/token [post]
func (c *Controller) Token(ctx *gin.Context) {
	var params schemas.OAuthLogin
	var user models.User

	if err := ctx.ShouldBind(&params); err == nil {
		c.DbProvider.Db.Where("username=?", strings.ToLower(params.Username)).First(&user)
	}

	if user.ID != 0 && user.Password == params.Password {
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
