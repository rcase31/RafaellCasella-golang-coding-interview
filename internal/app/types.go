package app

import (
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info"`
}

type StatesResponse struct {
	BaseResponse
	States []models.State `json:"data"`
}

type ReportsResponse struct {
	BaseResponse
	Reports []models.Report `json:"data"`
}

type PaginatedResponse[T any] struct {
	Items   []T `json:"items"`
	Total   int `json:"total"`
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
}

type PaginatedReportsResponse PaginatedResponse[ReportsResponse]
