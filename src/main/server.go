package main

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/config"
	"github.com/kartikeya/product_catalog_DIY/src/main/controller"
	"github.com/kartikeya/product_catalog_DIY/src/main/database"
	"github.com/kartikeya/product_catalog_DIY/src/main/repository"
	"github.com/kartikeya/product_catalog_DIY/src/main/router"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"gorm.io/gorm"
)

func defineApis(httpRouter router.Router, userController controller.UserControllerInterface, productController controller.ProductControllerInterface) {
	httpRouter.POST("/users/create", userController.CreateUser)
	httpRouter.POST("/products/add", productController.AddProducts)
	httpRouter.GET("/products/recommended/{n}", productController.GetRecommendedProducts)
	httpRouter.GET("/products", productController.GetProducts)
	httpRouter.PUT("/product/buy", productController.BuyProduct)
	httpRouter.GET("/product/{id}", productController.GetProductById)

}

func initializeDependencies() (router.Router, controller.UserControllerInterface, controller.ProductControllerInterface) {

	var httpRouter router.Router = router.NewMuxRouter()
	var DB *gorm.DB = database.ConnectPostgresDatabase()
	var salesRepository repository.SalesRepositoryInterface = &repository.SalesRepository{DB: DB}
	var userRepository repository.UserRepositoryInterface = &repository.UserRepository{DB: DB}
	var userService service.UserServiceInterface = &service.UserService{UserRepository: userRepository}
	var userController controller.UserControllerInterface = &controller.UserController{UserService: userService}
	var productRepository repository.ProductRepositoryInterface = &repository.ProductRepository{DB: DB}
	var productService service.ProductServiceInterface = &service.ProductService{ProductRepository: productRepository, UserService: userService, SalesRepository: salesRepository}
	var productController controller.ProductControllerInterface = &controller.ProductController{ProductService: productService}

	return httpRouter, userController, productController
}

func main() {
	httpRouter, userController, productController := initializeDependencies()
	defineApis(httpRouter, userController, productController)
	httpRouter.SERVE(config.PORT)
}
