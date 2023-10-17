package services

import (
	"github.com/Lalu-Mahato/go-dhan-das/models"
	"github.com/Lalu-Mahato/go-dhan-das/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (ps *ProductService) GetProducts() ([]models.Product, error) {
	products, err := ps.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(product *models.Product) error {
	err := ps.productRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}
