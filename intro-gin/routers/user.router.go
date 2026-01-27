package routers

import (
	"intro-gin/controllers"
	"intro-gin/services"

	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup) {
	// Init Service
	service := services.NewUserService()
	authService := services.NewAuthService()

	// Init Controller
	userController := controllers.NewUserController(service, authService)

	product := rg.Group("/users")
	{
		product.POST("/", userController.Register)
		product.POST("/login", userController.Login)
	}
}
