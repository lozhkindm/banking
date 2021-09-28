package domain

import "github.com/lozhkindm/banking/errs"

type Customer struct {
	Id        string
	Name      string
	City      string
	Zipcode   string
	BirthDate string
	Status    string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppError)
}
