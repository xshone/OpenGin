package routes

import (
	"opengin/server/controllers"
	"opengin/server/db"
	"opengin/server/middlewares"
	"opengin/server/websocket"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	// Database provider
	dbProvider := db.NewDbProvider()

	// Controller
	c := controllers.NewController(dbProvider)

	baseGroup := engine.Group("v1")

	// WebSocket
	hub := websocket.GetHub()
	go hub.Run()
	wsGroup := baseGroup.Group("ws")
	wsGroup.GET("/chat", c.Chat)
	wsGroup.GET("/get_message", c.GetMessage)

	// OAuth
	oauthGroup := baseGroup.Group("oauth")
	oauthGroup.POST("/token", c.Token)

	// Register
	registerGroup := baseGroup.Group("register")
	registerGroup.POST("/register", c.Register)
	registerGroup.GET("/check_user", c.CheckUsername)

	// Test
	testGroup := baseGroup.Group("test")
	testGroup.Use(middlewares.Auth())
	testGroup.GET("/hello", c.Hello)
	testGroup.GET("/:key", c.PathTest)
	testGroup.POST("/publish", c.PublishMessage)
	testGroup.POST("/redis_set", c.RedisSet)
	testGroup.POST("/redis_get", c.RedisGet)
}
