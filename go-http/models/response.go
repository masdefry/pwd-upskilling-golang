package models

type APIResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}