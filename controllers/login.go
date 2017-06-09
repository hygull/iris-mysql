package controllers

import (
	"fmt"
	"gopkg.in/kataras/iris.v5"
	"log"
	"restiris/conf"
	// "github.com/kataras/iris"
	// "net/http"
	// "encoding/json"
	// "reflect"
	// "strconv"
	// "strings"
	// "net/http"
)

func Login(ctx *iris.Context) {
		fmt.Println("4 Starting connection")
		// db, err := sql.Open("mysql", "hygull:admin@67@tcp(127.0.0.1:3306)/practice_db?charset=utf8")

		db := conf.DbObj

		email := string(ctx.FormValue("email"))
		password := string(ctx.FormValue("password"))		
		fmt.Println("URL => ",ctx.URLParam("email"))

		// myMap := ctx.FormValues()
		// var keyStrMap	string
		// for k,_ := range myMap {
		// 	keyStrMap = k 
		// 	break
		// }

		// type AuthData struct{
		// 	Email string `json:"email"`
		// 	Password string `json:"password"`
		// }
		// keyMap := AuthData{}
		// err = json.Unmarshal([]byte(keyStrMap), &keyMap)
		// fmt.Println(keyMap, reflect.TypeOf(keyMap))


		// email := keyMap.Email
		// password := keyMap.Password

		// if email == "" || password == "" {
		// 	fmt.Println("Email & Password(OR)")
			
		// 		fmt.Println("Email & Password(AND)")
		// 		// myMap := ctx.FormValues()
		// 		var keyStrMap	string
		// 		for k,_ := range myMap {
		// 			keyStrMap = k 
		// 			break
		// 		}

		// 		type AuthData struct{
		// 			Email string `json:"email"`
		// 			Password string `json:"password"`
		// 		}
		// 		keyMap := AuthData{}
		// 		err = json.Unmarshal([]byte(keyStrMap), &keyMap)
		// 		fmt.Println(keyMap, reflect.TypeOf(keyMap))


		// 		email = keyMap.Email
		// 		password = keyMap.Password

			if email == "" || password == "" {
				fmt.Println("Blank email or password")
				ctx.JSON(iris.StatusBadRequest, map[string]interface{}{"message": "email & password are required fields", "status": 400})
				return
			}
		// } 

		fmt.Println("Got", email, " & ", password)

		var fetchedId int
		var fetchedEmail, fetchedPassword, fetchedFirstname, fetchedLastname string
		rows, err := db.Query("Select id, email, password, fname, lname from users where email='"+email+"' AND password='"+password+"';")

		if err != nil {
			log.Fatal(err)
			ctx.JSON(iris.StatusInternalServerError, map[string]interface{}{"message": "error while processing SELECT(login) query", "status": 500})
			return
		} 

		found := false
		for rows.Next(){
			rows.Scan(&fetchedId, &fetchedEmail, &fetchedPassword, &fetchedFirstname, &fetchedLastname)
			found = true
		}
		
		if !found {
			ctx.JSON(iris.StatusUnauthorized, map[string]interface{}{"message": "email or password is incorrect", "status": 401})
			return
		}

		ctx.JSON(iris.StatusOK, map[string]interface{}{"message": "user found", "id": fetchedId, "email": fetchedEmail, "fullname": fetchedFirstname+" "+fetchedLastname, "status": 200})
		fmt.Println("Query succesfully executed while UPDATING record of single user")

		// db.Close()	
}