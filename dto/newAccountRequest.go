package dto

import (
	"github.com/lozhkindm/banking-lib/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000")
	}

	if strings.ToLower(r.AccountType) != "checking" && strings.ToLower(r.AccountType) != "saving" {
		return errs.NewValidationError("Account type should be checking or saving")
	}

	return nil
}
