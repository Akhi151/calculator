package controller

import (
	"database/sql"

	"fmt"
)

type Testtable2 struct {
	id int `json:"id"`
}

func db() {

	fmt.Println("Learning how to connect to the db(MySQL)")

	//connecting db

	db, err := sql.Open("mysql", "root:<pass123>@tcp(127.0.0.1:3306)/test")

	if err != nil {

		panic(err.Error()) //why we are using this error function?

	}

	defer db.Close()

	fmt.Println("Success!")

	insert, err := db.Query("INSERT INTO testtable2 VALUES('23')")

	if err != nil {

		panic(err.Error())

	}

	defer insert.Close()
}
