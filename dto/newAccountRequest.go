package dto

import (
	"banking/err"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *err.AppError {
	if r.Amount <= 5000 {
		return err.NewValidationError("amount : Amount must be greater than 5000")
	}
	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return err.NewValidationError("account_type : Account type must be either savings or checking")
	}
	return nil
}
