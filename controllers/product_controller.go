package controllers

import (
	"net/http"

	"github.com/Lalu-Mahato/go-dhan-das/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

func NewUserController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) FindUsers(c *gin.Context) {
	users, err := pc.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, users)
}
