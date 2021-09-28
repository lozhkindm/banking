package app

import (
	"github.com/gorilla/mux"
	"github.com/lozhkindm/banking/domain"
	"github.com/lozhkindm/banking/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
