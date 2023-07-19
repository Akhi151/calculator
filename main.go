package main

import (
	"calculator/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	controller.DbConnect()
	//fmt.Println("Success code is:", res)

	fmt.Println("Server is started...")
	router := httprouter.New()

	// Subtraction route
	router.GET("/sub", controller.Sub)
	//Addition route
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

	router.GET("/", welcome)

	// Start HTTP Listener
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Printf("HTTP Server stopped - %s", err)
	}

}

func welcome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Fprintf(w, "homepage")
}
