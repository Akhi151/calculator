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
<<<<<<< HEAD
	router.GET("/multiple", controller.Multiple)
=======
	router.GET("/div", controller.Div)
>>>>>>> a51e848ed2e5534bb83c0bf0dc54f42238df53d2

	// Start HTTP Listener
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}
}
