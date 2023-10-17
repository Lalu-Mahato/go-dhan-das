package repositories

import (
	"github.com/Lalu-Mahato/go-dhan-das/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(product *models.Product) error {
	err := pr.db.Create(product).Error
	return err
}

func (pr *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := pr.db.Find(&products).Error
	return products, err
}
