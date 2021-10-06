package domain

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lozhkindm/banking/errs"
	"github.com/lozhkindm/banking/logger"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) FindById(id string) (*Account, *errs.AppError) {
	var acc Account

	findSql := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"

	err := d.client.Get(&acc, findSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		} else {
			logger.Error("Error while scanning account: " + err.Error())
			return nil, errs.NewDatabaseError()
		}
	}

	return &acc, nil
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

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()

	if err != nil {
		logger.Error("Error while starting a new transaction: " + err.Error())
		return nil, errs.NewDatabaseError()
	}

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"

	result, _ := tx.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	var sqlAmount string

	if t.IsWithdrawal() {
		sqlAmount = "UPDATE accounts SET amount = amount - ? where account_id = ?"
	} else {
		sqlAmount = "UPDATE accounts SET amount = amount + ? where account_id = ?"
	}

	_, err = tx.Exec(sqlAmount, t.Amount, t.AccountId)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		logger.Error("Error while saving a transaction: " + err.Error())
		return nil, errs.NewDatabaseError()
	}

	err = tx.Commit()

	if err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		logger.Error("Error while committing a transaction: " + err.Error())
		return nil, errs.NewDatabaseError()
	}

	tId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting the last insert id for new a transaction: " + err.Error())
		return nil, errs.NewDatabaseError()
	}

	acc, e := d.FindById(t.AccountId)

	if e != nil {
		return nil, e
	}

	t.TransactionId = strconv.FormatInt(tId, 10)
	t.Amount = acc.Amount

	return &t, nil
}

func NewAccountRepositoryDB(client *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: client}
}
