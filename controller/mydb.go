package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() {
	fmt.Println("to connect to the db(MySQL)")
	//connecting db
	db, err := sql.Open("mysql", "root:Trident@1289@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//checking the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to mysql!!!")

	//Create a table

	createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            age INT
        )
    `

	_, err = db.Exec(createTableQuery)

	if err != nil {

		fmt.Println("Failed to create table:", err)

		return

	}
	fmt.Println("Table created successfully!")
	//Insert values into table

	insertQuery := "INSERT INTO users (name, age) VALUES (?, ?)"

	_, err = db.Exec(insertQuery, "Rohit", 20)

	if err != nil {

		fmt.Println("Failed to insert data:", err)

		return

	}

}
