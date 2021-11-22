package domain

import (
	"banking/err"
	"banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *err.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, errs := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if errs != nil {
		logger.Error("Error while creating a new account: " + errs.Error())
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}

	id, errs := result.LastInsertId()
	if errs != nil {
		logger.Error("Error while getting last inserted id for the new account: " + errs.Error())
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
