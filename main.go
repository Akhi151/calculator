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
	//Addition route
	router.GET("/add", controller.Add)

	router.GET("/", welcome)

	// Start HTTP Listener
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Fprintf(w, "homepage")
}
