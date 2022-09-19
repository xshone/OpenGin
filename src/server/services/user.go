package services

import (
	models "opengin/server/models"
	"opengin/server/schemas"
	"opengin/server/utils"
	"strings"
)

func CheckUsername(username string) bool {
	user := GetUser(username)

	return user.ID != 0
}

func GetUser(username string) models.User {
	var user models.User
	models.GetDB().Where("username = ?", strings.ToLower(username)).First(&user)

	return user
}

func CreateUser(params schemas.Register) string {
	passwordSalt := utils.GetUUID()
	hashedPassword := utils.HashPassword(params.Password, passwordSalt)

	user := models.User{
		Username:     params.Username,
		PasswordHash: hashedPassword,
		PasswordSalt: passwordSalt,
		Email:        params.Email,
	}

	result := models.GetDB().Create(&user)

	if result.Error == nil {
		token, err := utils.CreateToken(user.ID)

		if err == nil {
			return token
		}
	}

	return ""
}
