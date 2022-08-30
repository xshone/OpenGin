package main

import "opengin/server"

// @Version 1.0.0
// @Title OpenGin
// @Description Golang WebAPI which powered by OpenAPI 3.0 & Gin
// @ContactName Shone
// @ContactEmail xxxxxx@gmail.com
// @LicenseName MIT License
// @LicenseUrl https://mit-license.org
// @Server http://127.0.0.1 LocalTest
// @Server http://www.domainname.com Production
// @SecurityScheme OAuth2PasswordBearer oauth2ResourceOwnerCredentials /v1/oauth/token
func main() {
	server.InitAndStartServer()
}
