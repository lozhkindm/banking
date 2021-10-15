package domain

import (
	"github.com/lozhkindm/banking-lib/errs"
	"github.com/lozhkindm/banking/dto"
)

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}

type Customer struct {
	Id        string `db:"customer_id"`
	Name      string `db:"name"`
	City      string `db:"city"`
	Zipcode   string `db:"zipcode"`
	BirthDate string `db:"date_of_birth"`
	Status    string `db:"status"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:        c.Id,
		Name:      c.Name,
		City:      c.City,
		Zipcode:   c.Zipcode,
		BirthDate: c.BirthDate,
		Status:    c.statusAsText(),
	}
}

func (c Customer) statusAsText() string {
	status := "inactive"

	if c.Status == "1" {
		status = "active"
	}

	return status
}
