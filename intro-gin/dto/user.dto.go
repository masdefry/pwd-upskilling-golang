package dto

import (
	"intro-gin/models"
)

type UserDTO struct {
	Username string   `json:"username" binding:"required,min=3,max=100"`
	Email    string   `json:"email" binding:"required,email,max=255"`
	Password string   `json:"password" binding:"required,min=8"`
	Role     models.UserRole `json:"role" binding:"omitempty,oneof=ADMIN USER"`
}
