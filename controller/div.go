package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Values struct {
	Num1 int `json:"First"`
	Num2 int `json:"Second"`
}
type Responsee struct {
	Soln int `json:"res"`
}

func Div(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var numm Values
	err := json.NewDecoder(r.Body).Decode(&numm)
	if err != nil {
		fmt.Println("Error in decoding JSON")
	}
	res := division(numm.Num1, numm.Num2)
	response := Responsee{Soln: res}
	user, _ := json.Marshal(response)
	w.Write(user)
}

func division(a, b int) int {
	div := a / b
	return div
}
