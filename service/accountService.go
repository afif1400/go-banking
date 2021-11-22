package service

import (
	"banking/domain"
	"banking/dto"
	"banking/err"
	"time"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *err.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *err.AppError)
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *err.AppError) {
	//incoming request validation
	errs := req.Validate()
	if errs != nil {
		return nil, errs
	}

	//server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		// fmt.Println(req.AccountId)
		account, errs := s.repo.FindBy(req.AccountId)
		if errs != nil {
			return nil, errs
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, err.NewValidationError("Insufficient balance in the account")
		}
	}

	//if all is well, build the domain object and save the transaction
	t := domain.Transaction{
		// TransactionId: "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, errs := s.repo.SaveTransaction(t)
	if errs != nil {
		return nil, errs
	}

	response := transaction.ToTransactionResponseDto()

	return &response, nil

}
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
