package main

import (
	"encoding/json"

	"fmt"

	"log"

	"net/http"
)

type AddNum struct {
	Val1 int `json:"value1"`

	Val2 int `json:"value2"`
}

type Sum struct {
	Sum int `json:"sum"`
}

func main() {

	fmt.Println("Connection Established")

	handleRequests()

}

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "homepage")

}

func handleRequests() {

	http.HandleFunc("/", homepage)

	http.HandleFunc("/add", AdditionHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func AdditionHandler(w http.ResponseWriter, r *http.Request) {

	var addNum AddNum

	if err := json.NewDecoder(r.Body).Decode(&addNum); err != nil {

		http.Error(w, "Invalid Request", http.StatusBadRequest)

		return

	}

	sum := calculateSum(addNum.Val1, addNum.Val2)

	result := Sum{Sum: sum}

	response, err := json.Marshal(result)

	if err != nil {

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return

	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)

}

func calculateSum(addNum1, addNum2 int) int {

	sum := 0

	sum = addNum1 + addNum2

	return sum

}
