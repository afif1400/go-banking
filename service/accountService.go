package service

import (
	"banking/domain"
	"banking/dto"
	"banking/err"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *err.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *err.AppError) {

	errs := req.Validate()
	if errs != nil {
		return nil, errs
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, errs := s.repo.Save(a)

	if errs != nil {
		return nil, errs
	}

	reponseDto := newAccount.ToNewAccountResponseDto()

	return &reponseDto, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
