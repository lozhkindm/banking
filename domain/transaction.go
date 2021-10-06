package domain

import "github.com/lozhkindm/banking/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_Id"`
	AccountId       string  `db:"account_Id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToResponseDTO() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
