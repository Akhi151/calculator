package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
type Numbers struct{
	First int `json:"First"`
	Second int `json:"Second"`
}
type Response struct{
	Result int `json:"res"`
}

func homepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,  "homepage")
}

func handleRequests() {
	http.HandleFunc("/",homepage)
	http.HandleFunc("/sub",sub)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func sub(w http.ResponseWriter, r *http.Request){
	
	var num Numbers
	err:=json.NewDecoder(r.Body).Decode(&num)
	if err!=nil{
		fmt.Println("Error in decoding JSON")
	}
	res:=subtraction(num.First,num.Second)
	response:=Response{Result:res}
	user,_:=json.Marshal(response)
	w.Write(user)
	
}
func subtraction(a,b int) int{
	sub:=a-b
	return sub
}

func main() {
	fmt.Println("hhhh")
	handleRequests()
}

