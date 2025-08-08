package models

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	ID   string `json: "id"`
	Code string `json: "code"`
	Name string `json: "name"`
}

type Report struct {
	gorm.Model
	Num         int    `json: "num"`
	Header      string `json: "header"`
	Description string `json: "description"`
	Terms       string `json: "terms"`
}
