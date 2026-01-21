package services

import (
	"intro-gin/config"
	"intro-gin/models"
)

func GetProducts() []models.Product {
	var products []models.Product 
	
	result := config.DB.Find(&products)
	panic(result.Error)

	return products
}