package domain

import "github.com/lozhkindm/banking/errs"

type Customer struct {
	Id        string `json:"id" db:"customer_id"`
	Name      string `json:"name" db:"name"`
	City      string `json:"city" db:"city"`
	Zipcode   string `json:"zip_code" db:"zipcode"`
	BirthDate string `json:"birth_date" db:"date_of_birth"`
	Status    string `json:"status" db:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
