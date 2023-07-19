package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Number struct {
	Firsts  int `json:"First"`
	Seconds int `json:"Second"`
}
type Responses struct {
	Results int `json:"res"`
}

func Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var nums Number
	err := json.NewDecoder(r.Body).Decode(&nums)
	if err != nil {
		fmt.Println("Error in decoding JSON")
	}
	res := addition(nums.Firsts, nums.Seconds)
	response := Responses{Results: res}
	user, _ := json.Marshal(response)
	w.Write(user)

}

func addition(a, b int) int {
	sum := a + b
	return sum
}
