package controllers

import (
	"net/http"
	schemas "opengin/server/schemas"
	"opengin/server/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Title 注册新用户
// @Tag Register
// @Description - 用户名\n- 密码\n- 邮箱
// @Param  request body schemas.Register true "_"
// @Success  200 object schemas.UniResponse ""
// @Route /v1/register/register [post]
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

// @Title 检查用户名是否已注册
// @Tag Register
// @Description 检查用户名
// @Param username query string true "用户名"
// @Success  200 object schemas.UniResponse ""
// @Route /v1/register/check_user [get]
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
