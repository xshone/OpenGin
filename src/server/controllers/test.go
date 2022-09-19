package controllers

import (
	"net/http"
	schemas "opengin/server/schemas"
	"opengin/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @Title Hash Password
// @Tag Test
// @Description Hash Password
// @Security OAuth2PasswordBearer read write
// @Param text query string true "text here"
// @Route /v1/test/hash_password [get]
func (c *Controller) HashPassword(ctx *gin.Context) {
	var params schemas.TestText

	if err := ctx.ShouldBind(&params); err == nil {
		hashedPassword := utils.HashPassword(params.Text, utils.GetUUID())

		ctx.JSON(http.StatusOK, schemas.UniResponse{
			Code:    200,
			Message: hashedPassword,
			Data:    nil,
		})
		return
	}
}

// @Title Path Test
// @Tag Test
// @Description Test path parameters
// @Security OAuth2PasswordBearer read write
// @Param key path string true "Path key"
// @Route /v1/test/{key} [get]
func (c *Controller) PathTest(ctx *gin.Context) {
	key := ctx.Param("key")

	ctx.JSON(http.StatusOK, schemas.UniResponse{
		Code:    200,
		Message: "Hello " + key,
		Data:    nil,
	})
}

// @Title Publish message
// @Tag Test
// @Description Publish message here
// @Security OAuth2PasswordBearer read write
// @Param message query string true "Message"
// @Route /v1/test/publish [post]
func (c *Controller) PublishMessage(ctx *gin.Context) {
	message := ctx.Query("message")

	mq := utils.NewRabbitMqHandler(&utils.RabbitMqOptions{
		ExchangeName: "gin",
		QueueName:    "test",
		RoutingKey:   "",
	})
	defer mq.Close()

	mq.Publish(message)

	ctx.JSON(http.StatusOK, schemas.UniResponse{
		Code:    200,
		Message: "Published",
		Data:    nil,
	})
}

// @Title Set Message
// @Tag Test
// @Description Redis: Set message
// @Security OAuth2PasswordBearer read write
// @Param key query string true "Key"
// @Param value query string true "Message"
// @Route /v1/test/redis_set [post]
func (c *Controller) RedisSet(ctx *gin.Context) {
	key := ctx.Query("key")
	value := ctx.Query("value")

	redisHandler := utils.NewRedisHandler()

	redisHandler.Set(key, value, time.Hour)

	ctx.JSON(http.StatusOK, schemas.UniResponse{
		Code:    200,
		Message: "Done",
		Data:    nil,
	})
}

// @Title Get Message
// @Tag Test
// @Description Redis: Get message
// @Security OAuth2PasswordBearer read write
// @Param key query string true "Key"
// @Route /v1/test/redis_get [post]
func (c *Controller) RedisGet(ctx *gin.Context) {
	key := ctx.Query("key")

	redisHandler := utils.NewRedisHandler()

	value, ok := redisHandler.Get(key).(string)

	if ok {
		ctx.JSON(http.StatusOK, schemas.UniResponse{
			Code:    200,
			Message: value,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, schemas.UniResponse{
		Code:    200,
		Message: "",
		Data:    nil,
	})
}
