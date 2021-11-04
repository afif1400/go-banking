package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello nice to meet you")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Customers := []Customer{
		{Name: "Afif", City: "bengaluru", Zipcode: "560085"},
		{Name: "Ahmed", City: "Raichur", Zipcode: "584101"},
	}

	json.NewEncoder(w).Encode(Customers)
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe(":3000", nil)
}
