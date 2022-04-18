package main

import (
	"github.com/kartikeya/product_catalog_DIY/Database"
	"github.com/kartikeya/product_catalog_DIY/controller"
	"github.com/kartikeya/product_catalog_DIY/repository"
	"github.com/kartikeya/product_catalog_DIY/router"
	"github.com/kartikeya/product_catalog_DIY/service"
)

var (
	httpRouter        router.Router                = router.NewMuxRouter()
	productRepository repository.ProductRepository = repository.NewProductRepository(Database.ConnectPostgresDatabase())
	productService    service.ProductService       = service.NewProductService(productRepository)
	productController controller.ProductController = controller.NewProductController(productService)
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
