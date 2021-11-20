package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enviromnent variable not defined")
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getCustomersByStatus).Methods("GET").Queries("status", "{status}")
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)

	if err != nil {
		panic(err)
	}
}
