package main

import (
	"fmt"
	"intro-gin/config"
	"intro-gin/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := 8080;

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	config.ConnectDatabase()

	r := routers.SetupRouter()
	r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
			"message": "pong",
			})
	})

	r.Run(fmt.Sprintf(":%d", port));
}