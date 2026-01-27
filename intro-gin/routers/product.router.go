package routers

import (
	"intro-gin/controllers"
	"intro-gin/middlewares"
	"intro-gin/services"
	"os"

	"github.com/gin-gonic/gin"
)

func ProductRouter(rg *gin.RouterGroup) {
	// Init Service
	productService := services.NewProductService()

	// Init Controller
	productController := controllers.NewProductController(productService)

	product := rg.Group("/products")
	{
		product.GET("/", productController.GetAll)
		product.POST("/", middlewares.JWTAuth(os.Getenv("JWT_AUTH_SECRET_KEY")), middlewares.AuthorizeRole("ADMIN"), productController.Create)
	}
}
