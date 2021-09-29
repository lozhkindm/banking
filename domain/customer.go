package domain

import "github.com/lozhkindm/banking/errs"

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Zipcode   string `json:"zip_code"`
	BirthDate string `json:"birth_date"`
	Status    string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
