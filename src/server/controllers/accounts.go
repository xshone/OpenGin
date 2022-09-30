package controllers

import (
	"net/http"
	models "opengin/server/models"
	schemas "opengin/server/schemas"
	"opengin/server/services"
	"opengin/server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Title 检查用户名是否已注册
// @Tag Accounts
// @Description 检查用户名
// @Param username query string true "用户名"
// @Success  200 object schemas.UniResponse ""
// @Route /v1/accounts/check_user [get]
func (c *Controller) CheckUsername(ctx *gin.Context) {
	username := ctx.Query("username")

	if strings.Trim(username, " ") == "" {
		ctx.JSON(http.StatusInternalServerError, schemas.UniResponse{
			Code:    500,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}

	isExist := services.CheckUsername(username)

	if isExist {
		ctx.JSON(http.StatusInternalServerError, schemas.UniResponse{
			Code:    500,
			Message: "用户名已存在",
			Data:    true,
		})
	} else {
		ctx.JSON(http.StatusOK, schemas.UniResponse{
			Code:    200,
			Message: "用户名未注册",
			Data:    false,
		})
	}
}

// @Title 注册新用户
// @Tag Accounts
// @Description - 用户名\n- 密码\n- 邮箱
// @Param request body schemas.Register true "_"
// @Success 200 object schemas.UniResponse ""
// @Route /v1/accounts/register [post]
func (c *Controller) Register(ctx *gin.Context) {
	var params schemas.Register
	var message string
	var code int
	var token string

	status := http.StatusInternalServerError

	if err := ctx.ShouldBind(&params); err == nil {
		if params.Username != "" && params.Password != "" {
			token = services.CreateUser(params)

			if token != "" {
				status = http.StatusOK
				message = "注册成功"
				code = 200
			}
		}
	} else {
		message = err.Error()
	}

	ctx.JSON(status, schemas.UniResponse{
		Code:    code,
		Message: message,
		Data:    token,
	})
}

// @Title 登录
// @Tag Accounts
// @Description login
// @Param request body schemas.Login true "_"
// @Route /v1/accounts/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	var params schemas.Login
	var user models.User

	if err := ctx.ShouldBind(&params); err == nil {
		models.GetDB().Where("username=?", strings.ToLower(params.Username)).First(&user)
	}

	if user.ID != 0 && user.PasswordHash == utils.HashPassword(params.Password, user.PasswordSalt) {
		token, err := utils.CreateToken(user.ID)

		if err == nil {
			var data struct {
				UserId int    `json:"user_id"`
				Token  string `json:"token"`
			}
			data.UserId = user.ID
			data.Token = token

			ctx.JSON(http.StatusOK, schemas.UniResponse{
				Code:    200,
				Message: "Login Success",
				Data:    data,
			})
			return
		}
	}

	// http.Error(ctx.Writer, "Failed to login", http.StatusUnauthorized)
	ctx.JSON(http.StatusOK, schemas.UniResponse{
		Code:    403,
		Message: "Failed to login",
		Data:    nil,
	})
}
