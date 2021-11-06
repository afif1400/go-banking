package app

import (
	"banking/domain"
	"banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getCustomersByStatus).Methods("GET").Queries("status", "{status}")
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}
}
