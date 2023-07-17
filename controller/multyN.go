package controller
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Input struct{
	Input []int 	`json:"input"`
}

type Res struct{
	Result int `json:"result"`
}

func Multiple(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	var num Input
	err:=json.NewDecoder(r.Body).Decode(&num)
	if err!=nil{
		fmt.Println("Error in decoding JSON")
	}

	mul:=multiply1(num.Input)
	res:=Res{Result: mul}
	response,_:=json.Marshal(res)
	w.Header().Set("Content-Type","application/json")
	w.Write(response)
}

func multiply1(a []int) int{
	mul:=1
	for _,i:= range a{
		mul=mul* i
	}
	return mul
}