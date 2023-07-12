package main

import (
	"calculator/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
type Numbers struct{
	First int `json:"First"`
	Second int `json:"Second"`
}
type Response struct{
	Result int `json:"res"`
}

func main() {

	fmt.Println("Server is started...")
	router := httprouter.New()

	// Subtraction route
	router.GET("/sub", controller.Sub)
	router.GET("/mulN", controller.Multiple)

	// Start HTTP Listener
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}

