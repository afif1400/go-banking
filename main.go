package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello nice to meet you")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	Customers := []Customer{
		{Name: "Afif", City: "bengaluru", Zipcode: "560085"},
		{Name: "Ahmed", City: "Raichur", Zipcode: "584101"},
	}
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)
	} else {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customers)
	}
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe(":3000", nil)
}
