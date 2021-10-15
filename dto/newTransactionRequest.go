package dto

import "github.com/lozhkindm/banking-lib/errs"

const (
	DEPOSIT    = "deposit"
	WITHDRAWAL = "withdrawal"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != DEPOSIT && r.TransactionType != WITHDRAWAL {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}

func (r NewTransactionRequest) IsWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}
