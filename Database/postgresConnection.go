package Database

import (
	"fmt"
	"github.com/kartikeya/product_catalog_DIY/Model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//type DbObject interface {
//	First(*Model.Product, int)
//	AutoMigrate(*Model.Product)
//	Save(*Model.Product)
//	Find(*Model.Product)
//	Create(*Model.Product)
//}

var DB *gorm.DB
var err error

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=kartikeya dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//var DB1 *gorm.DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//DB = DB1
	fmt.Println(DB)
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Model.Product{})
}
