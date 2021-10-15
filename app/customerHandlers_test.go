package app

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/lozhkindm/banking-lib/errs"
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers
var mock *service.MockCustomerService

func setup(t *testing.T) func() {
	mock = service.NewMockCustomerService(gomock.NewController(t))
	ch = CustomerHandlers{service: mock}
	router = mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers)

	return func() {
		router = nil
	}
}

func TestGetAllCustomersReturns200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockResponse := []dto.CustomerResponse{
		{
			Id:        "100",
			Name:      "One hundred bucks",
			City:      "Los Angeles",
			Zipcode:   "110011",
			BirthDate: "1997-01-01",
			Status:    "1",
		},
		{
			Id:        "300",
			Name:      "Three hundred bucks",
			City:      "New York",
			Zipcode:   "220022",
			BirthDate: "1997-02-02",
			Status:    "1",
		},
	}

	mock.EXPECT().GetAllCustomers().Return(mockResponse, nil)

	r, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", http.StatusOK, w.Code))
	}
}

func TestGetAllCustomersReturns500(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mock.EXPECT().GetAllCustomers().Return(nil, errs.NewDatabaseError())

	r, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code))
	}
}
