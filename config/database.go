package config

import (
	"github.com/Lalu-Mahato/go-dhan-das/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
	)
}

var DB *gorm.DB

func DatabaseConnection() {
	dsn := "host=localhost user=postgres password=postgres dbname=dhan_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DatabaseMigration(db)
	DB = db
}
