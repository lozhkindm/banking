package service

import (
	"github.com/golang/mock/gomock"
	domain2 "github.com/lozhkindm/banking/domain"
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/errs"
	"github.com/lozhkindm/banking/mocks/domain"
	"testing"
	"time"
)

var service AccountService
var mock *domain.MockAccountRepository

func setup(t *testing.T) func() {
	mock = domain.NewMockAccountRepository(gomock.NewController(t))
	service = NewAccountService(mock)

	return func() {
		service = nil
	}
}

func TestNewAccountReturnsValidationError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	_, err := service.NewAccount(req)

	if err == nil {
		t.Error("Expected not nil error")
	}
}

func TestNewAccountReturnsError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	acc := domain2.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mock.EXPECT().Save(acc).Return(nil, errs.NewDatabaseError())

	_, err := service.NewAccount(req)

	if err == nil {
		t.Error("Expected not nil error")
	}
}

func TestNewAccountReturnsSuccess(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	acc := domain2.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	res := acc
	res.AccountId = "100"

	mock.EXPECT().Save(acc).Return(&res, nil)

	newAcc, err := service.NewAccount(req)

	if err != nil {
		t.Error("Expected nil error")
	}

	if newAcc.AccountId != res.AccountId {
		t.Error("Account ids are not equal")
	}
}
