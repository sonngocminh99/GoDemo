package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sonngocminh99/GoDemo/config"
	"github.com/sonngocminh99/GoDemo/controller"
	"github.com/sonngocminh99/GoDemo/repository"
	"github.com/sonngocminh99/GoDemo/service"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	productRepository repository.ProductRepository = repository.NewProductConnection(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	productRoutes := r.Group("api/products")
	{
		productRoutes.GET("/", productController.AllProducts)
		productRoutes.POST("/", productController.InsertProduct)
		productRoutes.GET("/:id", productController.FindProductById)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	r.Run()

}
