package calculator

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	//"strconv"
)

type Numbers struct {
	Values []int `json:"values"`
}

type Result struct {
	Sum int `json:"sum"`
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var numbers Numbers
	if err := json.NewDecoder(r.Body).Decode(&numbers); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sum := calculateSum(numbers.Values)
	result := Result{Sum: sum}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func calculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
