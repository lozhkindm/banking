package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/service"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	var req dto.NewAccountRequest

	req.CustomerId = mux.Vars(r)["customer_id"]

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		acc, err := h.service.NewAccount(req)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, acc)
		}
	}
}
