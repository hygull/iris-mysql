package controllers

import (
	"restiris/conf"
)

import (
	"fmt"
	"log"
	"gopkg.in/kataras/iris.v5"	
	"strconv"
)

func Product(ctx *iris.Context) {
		fmt.Println("2 Starting connection")

		db := conf.DbObj

		id := ctx.Param("id")

		_, err := strconv.Atoi(id)

		if err != nil {
			ctx.JSON(iris.StatusBadRequest, map[string]interface{}{"status": 400, "message" :"product id should be an integer"})
			return
		}

		rows, err := db.Query("select id, title, image, price, stock_status, target_url, merchant from products where id="+id)

		if err != nil {
			log.Fatal(err)
		}else{
			fmt.Println("Query succesfully excuted while fetching single product")
		}

		type Product struct {
			Id int 		 `json:"id"`
			Title string `json:"title"`
			Image string `json:"image"`
			Price float64  `json:"price"`
			StockStatus string `json:"stock_status"`
			TargetUrl string `json:"target_url"`
			Merchant string `json:"merchant"`
		}
			
		var product Product
	    found := false
		for rows.Next() {
			var id int
			var title, image, stoctkStatus, tagetUrl, merchant string
			var price float64

			fmt.Println("Scanning")
			rows.Scan(&id , &title, &image, &price, &stoctkStatus, &tagetUrl, &merchant)
			product = Product{id , title, image, price, stoctkStatus, tagetUrl, merchant}
			found = true
		}
		
		fmt.Println(found)

		if found==true {
			fmt.Println("Returning JSON representation of a product")
			ctx.JSON(iris.StatusOK, product)
		} else {
			fmt.Println("Returning JSON representation of a product")
			ctx.JSON(iris.StatusBadRequest, map[string]interface{}{"status": 400, "message" :"product with this id does not exist"})
		}
}