package app

import (
	"github.com/DeVasu/gotoken-api/controllers/cashiers"
	"github.com/DeVasu/gotoken-api/controllers/categories"
	"github.com/DeVasu/gotoken-api/controllers/ping"
	"github.com/DeVasu/gotoken-api/controllers/products"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/cashiers", cashiers.List)
	router.GET("/cashiers/:cashierId", cashiers.GetById)
	router.POST("/cashiers", cashiers.Create)
	router.PUT("/cashiers/:cashierId", cashiers.Update)
	router.DELETE("/cashiers/:cashierId", cashiers.Delete)
	router.POST("/cashiers/:cashierId/login", cashiers.Common)
	router.POST("/cashiers/:cashierId/logout", cashiers.Common)
	router.GET("/cashiers/:cashierId/passcode", cashiers.Common)

	router.GET("/categories", categories.List)
	router.POST("/categories", categories.Create)
	router.GET("/categories/:categoryId", categories.GetById)
	router.PUT("/categories/:categoryId", categories.Update)
	router.DELETE("/categories/:categoryId", categories.Delete)

	router.POST("/products", products.Create)
	router.GET("/products", products.List)
	router.GET("/products/:productId", products.GetById)

}
