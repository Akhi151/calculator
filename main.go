package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numbers struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Mult struct {
	Mult int `json:"mult"`
}

func main() {

	fmt.Println("hello world")
	handleRequests()
}

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "homepage")
}

func handleRequests() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/test", MultiplicationHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func MultiplicationHandler(w http.ResponseWriter, r *http.Request) {

	var multNum Numbers
	if err := json.NewDecoder(r.Body).Decode(&multNum); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	mult := calculateMultiple(multNum.Num1, multNum.Num2)

	result := Mult{Mult: mult}

	response, err := json.Marshal(result)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func calculateMultiple(multNum1, multNum2 int) int {

	mult := multNum1 * multNum2
	fmt.Println("Multiplication of a number ", multNum1, "*", multNum2, "is: ", mult)
	return mult

}
