package routes

import (
	"github.com/Lalu-Mahato/go-dhan-das/config"
	"github.com/Lalu-Mahato/go-dhan-das/controllers"
	"github.com/Lalu-Mahato/go-dhan-das/repositories"
	"github.com/Lalu-Mahato/go-dhan-das/services"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	db := config.DB
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewUserController(productService)

	api := router.Group("/api/v1")
	api.GET("/products", productController.FindUsers)
}
