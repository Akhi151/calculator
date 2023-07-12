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
	router.GET("/add", controller.Add)
	router.GET("/addn", controller.Addn)
	router.GET("/multiply", controller.Multiply)
	router.GET("/multiple", controller.Multiple)

	// Start HTTP Listener
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}
