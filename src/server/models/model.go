package models

import (
	"opengin/server/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbUser := config.Settings.Database.User
	dbPwd := config.Settings.Database.Password
	dbHost := config.Settings.Database.Host
	dbPort := config.Settings.Database.Port
	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/godb?charset=utf8mb4&parseTime=True&loc=Local"
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := gormDb.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(config.Settings.Database.ConnMaxLifetimeHours))
	DB = gormDb

	return gormDb
}

func GetDB() *gorm.DB {
	return DB
}

func Migrate() {
	DB.AutoMigrate(
		&User{},
		&Role{},
	)
}
