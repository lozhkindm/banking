package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/lozhkindm/banking/service"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "Leha", City: "Ejevsk", Zipcode: "100505"},
	//	{Name: "Lexa", City: "Moscow", Zipcode: "900303"},
	//}

	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		log.Fatal(err)
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
