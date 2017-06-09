package controllers


import (
	"fmt"
	"log"
	"gopkg.in/kataras/iris.v5"
	"restiris/conf"
	// "github.com/kataras/iris"
	// "net/http"
	// "encoding/json"
	// "reflect"
	// "strconv"
	// "strings"
	// "net/http"
)

func CreateProduct(ctx *iris.Context) {
		fmt.Println("3 Starting connection")
		// db, err := sql.Open("mysql", "hygull:admin@67@tcp(127.0.0.1:3306)/practice_db?charset=utf8")
		db := conf.DbObj

		// if err != nil {
		// 	log.Fatal(err, "H")
		// }

		// forms := ctx.FormValues()

		title := ctx.FormValue("title")
		image := ctx.FormValue("image") 
		price := ctx.FormValue("price")
		stockStatus := ctx.FormValue("stock_status")
		tagetUrl := ctx.FormValue("target_url")
		merchant := ctx.FormValue("merchant")
		// merchant := forms["merchant"][0]

		// fmt.Println("================")
		// for key, value := range forms {
		// 	// fmt.Println("key:", key, len(key),", value: ", value)

		// 	if strings.TrimSpace(key) == "merchant" {
		// 		// fmt.Println ("Fetching .....")
		// 		// fmt.Println(reflect.TypeOf(value))
		// 		merchant = value[0]
		// 	}
		// }

		// fmt.Println(".....", forms["merchant"])
		// fmt.Println()
		// fmt.Println("================")

		fmt.Println("Got", title, image, price, stockStatus)
		fmt.Println("Target: ", tagetUrl, "Merchant: ",merchant)

		query := "INSERT INTO products(title, image, price, stock_status, target_url, merchant) values('"+
					  string(title) +"', '"+string(image)+"', "+string(price)+", "+string(stockStatus)+", '"+string(tagetUrl)+"', '"+string(merchant)+"')" 
		
		fmt.Println("Executing ", query)
		stmt, err := db.Prepare(query)

		if err != nil {
			ctx.JSON(iris.StatusInternalServerError, map[string]interface{}{"message": "error while processing INSERT query, all the fields are required","required_fields": "title, image, price, stock_status, target_url" ,"status": 500})
			return
		} 

		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(iris.StatusOK, map[string]interface{}{"message": "product successfully created", "status": 200})
			fmt.Println("Query succesfully excuted while INSERTING record of single product")
		}
		
		// db.Close()	
}