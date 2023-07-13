package main

import (
	"calculator/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("Server is started...")
	router := httprouter.New()

	// Subtraction route
	router.GET("/sub", controller.Sub)
	// Add route
	router.GET("/add", controller.Add)
	// Adding 'n' number route
	router.GET("/addn", controller.Addn)
	// multipy route
	router.GET("/multiply", controller.Multiply)
	// multiply 'n'  route
	router.GET("/multiple", controller.Multiple)
	// div route
	router.GET("/div", controller.Div)

	// Start HTTP Listener
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}
