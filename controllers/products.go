package controllers

import (
	"fmt"
	"gopkg.in/kataras/iris.v5"
)

import (
	"restiris/conf"
)

func Products(ctx *iris.Context) {
		fmt.Println("1 Starting connection")
		// db, err := sql.Open("mysql", "<username>:<password>@tcp(127.0.0.1:<port>)/<dbname>"	)
		// db, err := sql.Open("mysql", "hygull:admin@67@tcp(127.0.0.1:3306)/practice_db?charset=utf8")
		db := conf.DbObj
		rows, err := db.Query("select id, title, image, price, stock_status, target_url, merchant from products")

		if err != nil {
			ctx.JSON(iris.StatusInternalServerError, map[string]interface{}{"message": "error while query execution"})
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
			
		var products []Product

		for rows.Next() {
			var id int
			var title, image, stoctkStatus, tagetUrl, merchant string
			var price float64

			rows.Scan(&id , &title, &image, &price, &stoctkStatus, &tagetUrl, &merchant)
			products = append(products, Product{id , title, image, price, stoctkStatus, tagetUrl, merchant})
		}
		
		// productsBytes, _ := json.Marshal(&products)
		
		// w.Write(productsBytes)

		fmt.Println("Returning JSON representation of all products")
		ctx.JSON(iris.StatusOK, products)
		db.Close()
}
