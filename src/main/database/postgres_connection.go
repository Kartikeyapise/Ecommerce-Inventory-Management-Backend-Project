package database

import (
	"fmt"
	"github.com/kartikeya/product_catalog_DIY/src/main/config"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectPostgresDatabase() *gorm.DB {
	dsn := config.DATABASE_URL
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	AutoMigrateGormModels(DB)
	log.Println("Postgres database connected.......")
	return DB
}

func AutoMigrateGormModels(DB *gorm.DB) {
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Sales{})
}
