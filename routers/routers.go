package routers

import "gopkg.in/kataras/iris.v5"
import "fmt"

import (
	"restiris/conf"
	"restiris/controllers"
)

func Route() {

	/* To get all the products */
	iris.Get("/products/", controllers.Products)

	// iris.Get("/products/:id", Product)
	// iris.Post("/products/create/", CreateProduct)
	// iris.Put("/products/update/", UpdateProduct)
	// iris.Post("/login/", Login)

	iris.Listen(":" + conf.ServerPort)

	fmt.Println("\nBye")
}
