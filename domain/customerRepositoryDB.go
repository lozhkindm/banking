package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lozhkindm/banking/errs"
	"github.com/lozhkindm/banking/logger"
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
		return nil, errs.NewDatabaseError()
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
			return nil, errs.NewDatabaseError()
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: client}
}
