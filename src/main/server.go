package main

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/Database"
	"github.com/kartikeya/product_catalog_DIY/src/main/controller"
	"github.com/kartikeya/product_catalog_DIY/src/main/repository"
	"github.com/kartikeya/product_catalog_DIY/src/main/router"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"gorm.io/gorm"
)

var (
	httpRouter        router.Router                         = router.NewMuxRouter()
	DB                *gorm.DB                              = Database.ConnectPostgresDatabase()
	productRepository repository.ProductRepositoryInterface = &repository.Repository{DB: DB}
	productService    service.ProductServiceInterface       = &service.Service{ProductRepository: productRepository}
	productController controller.ProductControllerInterface = &controller.Controller{ProductService: productService}
)

func defineApis() {
	httpRouter.POST("/products", productController.AddProducts)
	httpRouter.GET("/product/{id}", productController.GetProductById)
	httpRouter.GET("/products", productController.GetProducts)
	httpRouter.PUT("/buyProduct/{id}/{quantity}", productController.BuyProduct)
	httpRouter.GET("/getTop5Products", productController.GetTop5Products)
}

func main() {
	defineApis()
	const port string = ":9000"
	httpRouter.SERVE(port)
}
