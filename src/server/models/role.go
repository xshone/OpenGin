package models

import (
	base "opengin/server/models/core"
)

type Role struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	IsDefault   bool
	Time        base.CreateUpdateTime `gorm:"embedded"`
}

func (u *Role) TableName() string {
	return "role"
}
