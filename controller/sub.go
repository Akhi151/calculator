package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Numbers struct {
	First  int `json:"First"`
	Second int `json:"Second"`
}
type Response struct {
	Result int `json:"res"`
}

func Sub(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var num Numbers
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		fmt.Println("Error in decoding JSON")
	}
	res := Subtraction(num.First, num.Second)
	response := Response{Result: res}
	user, _ := json.Marshal(response)
	w.Write(user)
}

func Subtraction(a, b int) int {
	sub := a - b
	return sub
}
