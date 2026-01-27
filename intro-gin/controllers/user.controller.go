package controllers

import (
	"intro-gin/dto"
	"intro-gin/models"
	"intro-gin/services"
	"intro-gin/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
	AuthService *services.AuthService
}

func NewUserController(service *services.UserService, authService *services.AuthService) *UserController {
	return &UserController{Service: service, AuthService: authService}
}

func (ctr *UserController) Register(c *gin.Context) {
	var req dto.UserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": utils.ValidationErrorResponse(err),
			"data":    nil,
		})
		return
	}

	user, err := ctr.Service.Register(req.Username, req.Password, req.Role)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Register user successful",
		"data":    user,
	})
}



func (ctr *UserController) Login(c *gin.Context) {
	var req *models.User

	// 1️⃣ Bind & validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": true, 
			"message": err.Error(),
			"data": nil,
		})
		return
	}

	// 2️⃣ Call service login
	user, err := ctr.Service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false, 
			"message": "Invalid username or password",
			"data": nil,
		})
		return
	}

	// 3️⃣ Generate JWT
	token, err := ctr.AuthService.GenerateToken(user.Id.UUID.String(), user.Role, 24*time.Hour, os.Getenv("JWT_AUTH_SECRET_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": true, 
			"message": err.Error(),
			"data": nil,
		})
		return
	}

	// 4️⃣ Response (tanpa password)
	c.JSON(http.StatusOK, gin.H{
		"success": true, 
		"message": "Login account successful", 
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":       user.Id,
				"username": user.Username,
				"role":     user.Role,
			},
		},
	})
}
