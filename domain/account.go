package domain

import (
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/errs"
)

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/lozhkindm/banking/domain AccountRepository
type AccountRepository interface {
	FindById(id string) (*Account, *errs.AppError)
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
