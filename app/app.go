package app

import "net/http"

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		panic(err)
	}
}
