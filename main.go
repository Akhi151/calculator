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

	// Start HTTP Listener
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}
