package main

import (
	deliverygin "github.com/ducbm95/go-clean-arch-template/product/delivery/gin"
	"github.com/ducbm95/go-clean-arch-template/product/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// db config
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// service layer
	productService := service.NewProductService(db)

	// delivery layer
	productController := deliverygin.NewProductController(productService)

	// gin config
	r := gin.Default()
	r.POST("/product/create", productController.Create)
	r.Run()
}
