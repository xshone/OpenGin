package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"opengin/server/config"
	"opengin/server/models"
	router "opengin/server/routes"
	oas "opengin/server/swagger/oas"
	"opengin/server/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

const (
	ModeDebug   string = "debug"
	ModeRelease string = "release"
	ModeTest    string = "test"
)

type Application struct {
	GinEngine *gin.Engine
	DB        *gorm.DB
}

func loadSettings(configPath string) {
	jsonFile, err := os.Open(configPath)

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

func New(mode string) *Application {
	runMode := gin.DebugMode

	if mode == ModeRelease {
		runMode = gin.ReleaseMode
	} else if mode == ModeTest {
		runMode = gin.TestMode
	}

	gin.SetMode(runMode)
	ginEngine := gin.New()

	// Middleware
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}), gin.ErrorLogger())

	// Database
	database := models.Init()
	models.Migrate()

	// Router
	router.InitRoutes(ginEngine)

	// Swagger
	go oas.GenerateOpenApiSpec("server/swagger/docs/openapi.json", false, true)
	ginEngine.Static("/swagger", "./server/swagger/dist")
	ginEngine.Static("/docs", "./server/swagger/docs")

	app := new(Application)
	app.GinEngine = ginEngine
	app.DB = database

	return app
}

func InitConfig(mode string) {
	// Get current path
	currentPath := utils.GetCurrentPath()

	// Load config
	configPath := currentPath + "/server/config/settings_dev.json"

	if mode == ModeRelease {
		configPath = currentPath + "/server/config/settings.json"

	}

	loadSettings(configPath)
}

func (a *Application) Run() {
	server := &http.Server{
		Addr:    ":" + config.Settings.Server.Port,
		Handler: a.GinEngine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	fmt.Println("Server exiting")
}
