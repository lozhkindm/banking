package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/lozhkindm/banking/errs"
	"github.com/lozhkindm/banking/logger"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewDatabaseError()
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewDatabaseError()
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: client}
}
