package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}
}
