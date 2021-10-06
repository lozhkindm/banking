package service

import (
	"github.com/lozhkindm/banking/domain"
	"github.com/lozhkindm/banking/dto"
	"github.com/lozhkindm/banking/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	acc := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAcc, err := s.repo.Save(acc)

	if err != nil {
		return nil, err
	}

	res := newAcc.ToNewAccountResponseDto()

	return &res, nil
}

func (s DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.IsWithdrawal() {
		acc, err := s.repo.FindById(req.AccountId)

		if err != nil {
			return nil, err
		}

		if !acc.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient funds on the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	tr, err := s.repo.SaveTransaction(t)

	if err != nil {
		return nil, err
	}

	res := tr.ToResponseDTO()

	return &res, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
