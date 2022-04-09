package Database

import (
	"fmt"
	"github.com/kartikeya/product_catalog_DIY/Model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=kartikeya dbname=mytestdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Model.Product{})
}
