package service

import (
	"github.com/lozhkindm/banking/domain"
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/errs"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/lozhkindm/banking/service CustomerService
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var response []dto.CustomerResponse

	for _, v := range c {
		response = append(response, v.ToDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
