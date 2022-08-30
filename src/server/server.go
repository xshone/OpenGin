package server

import (
	"encoding/json"
	"io"
	"opengin/server/config"
	router "opengin/server/routes"
	oas "opengin/server/swagger/oas"
	"opengin/server/utils"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func InitAndStartServer() {
	env := os.Getenv("env")

	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	}

	ginEngine := gin.New()

	// Load config
	loadSettings()

	// Middleware
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}), gin.ErrorLogger())

	// Router
	router.InitRoutes(ginEngine)

	// Swagger
	go oas.GenerateOpenApiSpec("server/swagger/docs/openapi.json", false, true)
	ginEngine.Static("/swagger", "./server/swagger/dist")
	ginEngine.Static("/docs", "./server/swagger/docs")

	// Run
	ginEngine.Run(":" + config.Settings.Server.Port)
}

func loadSettings() {
	// Get current path
	currentPath := utils.GetCurrentPath()
	jsonFile, err := os.Open(currentPath + "/server/config/settings.json")

	if err != nil {
		panic(err)
	}
	jsonData, err := io.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &config.Settings)

	if err != nil {
		panic(err)
	}
}