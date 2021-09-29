package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lozhkindm/banking/errs"
	"github.com/lozhkindm/banking/logger"
	"time"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	findSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	err := d.client.Select(&customers, findSql)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	var customer Customer

	findSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	err := d.client.Get(&customer, findSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:13306)/banking")

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client: client}
}
