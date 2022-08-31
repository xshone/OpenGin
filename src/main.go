package main

import (
	"opengin/server"
	"os"
)

// @Version 1.0.0
// @Title OpenGin
// @Description Golang WebAPI which powered by OpenAPI 3.0 & Gin
// @ContactName Shone
// @ContactEmail xxxxxx@gmail.com
// @LicenseName MIT License
// @LicenseUrl https://mit-license.org
// @SecurityScheme OAuth2PasswordBearer oauth2ResourceOwnerCredentials /v1/oauth/token
func main() {
	runMode := os.Getenv("mode")
	server.InitConfig(runMode)
	app := server.New(runMode)
	app.Run()
}
