package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Val struct {
	FirstInp  int `json:"First"`
	SecondInp int `json:"Second"`
}
type Solution struct {
	Report int `json:"res"`
}

func Multiply(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var val Val
	err := json.NewDecoder(r.Body).Decode(&val)
	if err != nil {
		fmt.Println("Error in decoding JSON")
	}
	res := multiplication(val.FirstInp, val.SecondInp)
	response := Solution{Report: res}
	user, _ := json.Marshal(response)
	w.Write(user)
}

func multiplication(a, b int) int {
	prod := a * b
	return prod
}
