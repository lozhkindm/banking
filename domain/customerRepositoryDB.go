package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lozhkindm/banking/errs"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findSql)

	if err != nil {
		log.Println("Error while querying customer table", err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.BirthDate, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers", err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	findSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(findSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.BirthDate, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning customer", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:13306)/banking")

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client: client}
}
