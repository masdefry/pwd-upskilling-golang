package controllers

import (
	"fmt"
	"intro-gin/models"
	"intro-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (ctr *ProductController) GetAll(c *gin.Context) {
	products, err := ctr.Service.GetAll()

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
		"message": "Get products successful",
		"data":    products,
	})
}

func (ctr *ProductController) Create(c *gin.Context) {
	var req *models.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := ctr.Service.Create(req.Name, req.Price, req.Stock)

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
		"message": "Create product successful",
		"data":    nil,
	})
}

func (ctr *ProductController) Update(c *gin.Context) {
	id := c.Param("id")

	var req *models.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := ctr.Service.Create(req.Name, req.Price, req.Stock)

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
		"message": fmt.Sprint("Update product with id=%s successful", id),
		"data":    nil,
	})
}