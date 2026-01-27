package services

import (
	"intro-gin/config"
	"intro-gin/models"
)

type ProductService struct{}
func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAll()([]models.Product, error) {
	var products []models.Product 
	
	result := config.DB.Find(&products)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (s *ProductService) Create(name string, price int, stock *int) (error) {
	product := &models.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	if err := config.DB.Create(product).Error; err != nil {
		return err
	}

	return nil
}