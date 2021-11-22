package domain

import (
	"banking/dto"
)

const WITHDRAWL = "WITHDRAWL"
const DEPOSIT = "DEPOSIT"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawl() bool {
	return t.TransactionType == WITHDRAWL
}

// to transaction DTO
func (t Transaction) ToTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
