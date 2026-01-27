package routers

import (
	"intro-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.RecoverMiddleware())
	api := r.Group("/api")
	{
		ProductRouter(api)   // Load Product Router
		UserRouter(api) // Load User Router
	}

	return r
}
