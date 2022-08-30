package controllers

import (
	"opengin/server/db"
)

type Controller struct {
	DbProvider *db.DbProvider
}

func NewController(dp *db.DbProvider) *Controller {
	return &Controller{DbProvider: dp}
}

type Message struct {
	Message string `json:"message" example:"message"`
}
