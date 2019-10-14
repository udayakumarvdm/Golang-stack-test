package main

import (
	"fmt"

	usecase "stack/test"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//Connect the Postgres Database
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "postgres"
	dbName := "test"

	db, err := gorm.Open("postgres",
		"host="+dbHost+" port="+dbPort+" user="+dbUser+" dbname="+dbName+" password="+dbPass+" sslmode=disable")
	if err != nil {
		fmt.Println("Connection failed ==>", err)
	} else {
		fmt.Println("Test database connection Established successfully")
	}
	defer db.Close()

	//Call the bussiness logic
	req, err := usecase.Test()
	if err != nil {
		panic(err)
	}

	//Inser record into tables
	db.Create(req)

	fmt.Println(db.Error)

	forever := make(chan bool)

	go func() {
		fmt.Println("Test server start successfully")

	}()

	<-forever
}
