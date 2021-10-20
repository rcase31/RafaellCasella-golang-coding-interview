package models

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Id   string `json: "id"`
	Code string `json: "code"`
	Name string `json: "name"`
}
