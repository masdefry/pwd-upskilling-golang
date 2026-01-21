package routers

import (
	"github.com/gin-gonic/gin"
	"intro-gin/controllers"
)

func ProductRouter(rg *gin.RouterGroup) {
	product := rg.Group("/products")
	{
		product.GET("/", controllers.GetProducts)
	}
}
