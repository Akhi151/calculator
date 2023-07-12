package controller

type Number struct{
	Values []int `json:"values"`
}

type Res struct{
	Result int `json:"result"`
}

func Multiple(w http.ResponseWriter, r *http.Request){
	var num Number
	err:=json.NewDecoder(r.Body).Decode(&num)
	if err!=nil{
		fmt.Println("Error in decoding JSON")
	}

	mul:=multiply(num.Values)
	res:=Res{Result: mul}
	response,_:=json.Marshal(res)
	w.Header().Set("Content-Type","application/json")
	w.Write(response)
}

func multiply(a []int) int{
	mul:=1
	for _,i:= range a{
		mul*= i
	}
	return mul
}