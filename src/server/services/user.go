package services

import (
	models "opengin/server/models"
	"opengin/server/schemas"
	"opengin/server/utils"
	"strings"

	"gorm.io/gorm"
)

func CheckUsername(db *gorm.DB, username string) bool {
	user := GetUser(db, username)

	return user.ID != 0
}

func GetUser(db *gorm.DB, username string) models.User {
	var user models.User
	db.Where("username = ?", strings.ToLower(username)).First(&user)

	return user
}

func CreateUser(db *gorm.DB, params schemas.Register) string {
	user := models.User{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
	}

	result := db.Create(&user)

	if result.Error == nil {
		token, err := utils.CreateToken(user.ID)

		if err == nil {
			return token
		}
	}

	return ""
}
