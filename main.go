package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,  "homepage")
}

func handleRequests() {
	http.HandleFunc("/",homepage)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main() {
	fmt.Println("hhhh")
	handleRequests()
}
