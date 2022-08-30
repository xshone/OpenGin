package models

import (
	base "opengin/server/models/core"
)

type User struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Email    string
	Time     base.CreateUpdateTime `gorm:"embedded"`
}

func (u *User) TableName() string {
	return "user"
}
