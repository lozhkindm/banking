package domain

import (
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/errs"
)

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
}

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}
