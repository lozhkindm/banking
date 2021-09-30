package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lozhkindm/banking/config"
	"github.com/lozhkindm/banking/domain"
	"github.com/lozhkindm/banking/service"
	"log"
	"net/http"
	"time"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDbClient()

	// wiring
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient)),
	}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(config.NewServerConfig().AsString(), router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", config.NewDbConfig().AsDataSource())

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
