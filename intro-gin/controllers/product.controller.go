package controllers

import (
	"intro-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := services.GetProducts()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get products successful",
		"data":    products,
	})
}
