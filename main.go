/**
	{
		"created_on": "6 May 2017",
		"aim": "Rest API development for web, android & iOS application"
		"coded_by": "Rishikesh Agrawani",
	}
*/
package main

import (
	_"github.com/go-sql-driver/mysql"	
	"database/sql"
	"fmt"
)

import (
	"restiris/routers"
	"restiris/conf"
)

func main() {
	connString := conf.DbUser + ":" + conf.DbPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.Db + "?charset=utf8"
	fmt.Println("\n[SharpSeller] Opening a DB connection with ", connString)
	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("\n[SharpSeller] Error while db connection, wait we will fix it soon")
		return
	} 

	//If connection is Ok then assign the db value to global Db(to make it accessible in other packages)
	conf.DbObj = db

	routers.Route()
}