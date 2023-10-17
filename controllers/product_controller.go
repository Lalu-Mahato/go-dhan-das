package controllers

import (
	"net/http"
	"strings"

	"github.com/Lalu-Mahato/go-dhan-das/models"
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

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid product request")
		return
	}

	if strings.TrimSpace(product.Name) == "" {
		c.JSON(http.StatusBadRequest, "Invalid product name")
		return
	}

	if err := pc.productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create new product")
		return
	}
	c.JSON(http.StatusCreated, product)
}
