package Database

import (
	"fmt"
	"github.com/kartikeya/product_catalog_DIY/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectTestDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=kartikeya dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&entity.Product{})
	fmt.Println("Postgres Database connected.......")
	return DB
}
