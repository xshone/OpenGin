package db

import (
	"opengin/server/config"
	models "opengin/server/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbProvider struct {
	Db *gorm.DB
}

func NewDbProvider() *DbProvider {
	dbUser := config.Settings.Database.User
	dbPwd := config.Settings.Database.Password
	dbHost := config.Settings.Database.Host
	dbPort := config.Settings.Database.Port
	// dsn := "root:123456@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/godb?charset=utf8mb4&parseTime=True&loc=Local"
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	gormDb.AutoMigrate(&models.User{}, &models.Role{})

	sqlDB, err := gormDb.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(config.Settings.Database.ConnMaxLifetimeHours))

	return &DbProvider{Db: gormDb}
}
