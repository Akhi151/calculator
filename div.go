package main

import (

    "encoding/json"

    "fmt"

    "net/http"

)

type Numbers struct {

    First  int `json:"first"`

    Second int `json:"second"`

}

type Response struct {

    Result int `json:"divv"`

}

func main() {

    fmt.Println("hhh")

    http.HandleFunc("/div", division)

    http.ListenAndServe(":8080", nil)

}

func division(w http.ResponseWriter, r *http.Request) {

    var num Numbers

    err := json.NewDecoder(r.Body).Decode(&num)

    if err != nil {

        fmt.Println("Error in Decoding")

    }

    divv := divide(num.First, num.Second)

    res := Response{Result: divv}

    res1, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")

    w.Write(res1)
}

func divide(a, b int) int {

    div := a / b

    return div

}