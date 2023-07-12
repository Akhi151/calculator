package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Num struct {
	// One  int `json:"First"`
	// Two int `json:"Second"`
	Values []int `json:"values"`
}
type Resp struct {
	Resu int `json:"res"`
}

func Addn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var number Num
	err := json.NewDecoder(r.Body).Decode(&number)
	if err != nil {
		fmt.Println("Error in decoding JSON")
	}
	res := additionN(number.Values)
	response := Resp{Resu: res}
	user, _ := json.Marshal(response)
	w.Write(user)
}

func additionN(number []int) int {
	//addn := a - b
	addn := 0
	for _, nums := range number {
		addn += nums
	}
	return addn
}
