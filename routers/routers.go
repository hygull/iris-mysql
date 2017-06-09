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

	/* To get a single product */
	iris.Get("/products/:id", controllers.Product)

	/* To create a product */
	iris.Post("/products/create/", controllers.CreateProduct)

	// iris.Put("/products/update/", UpdateProduct)

	/* To login */
	iris.Post("/login/", controllers.Login)

	// iris.Post("/register/", Register)
	// iris.Get("/users/", Users)
	// iris.Get("/users/:id", User)

	iris.Listen(":" + conf.ServerPort)

	fmt.Println("\nBye")
}
