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

// transaction = make an entry in the transaction table and update the balance of the accounts table
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *err.AppError) {
	tx, errs := d.client.Begin()
	if errs != nil {
		logger.Error("Error while beginning transaction: " + errs.Error())
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}

	sqlInsert := "INSERT INTO transactions (account_id, transaction_type, amount, transaction_date) values (?,?,?,?)"
	result, _ := tx.Exec(sqlInsert, t.AccountId, t.TransactionType, t.Amount, t.TransactionDate)

	//updating account table
	sqlUpdate := "UPDATE accounts SET amount = ? WHERE account_id = ?"
	if t.IsWithdrawl() {
		sqlUpdate = "UPDATE accounts SET amount = amount - ? WHERE account_id = ?"
	} else {
		sqlUpdate = "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	}

	_, errs = tx.Exec(sqlUpdate, t.Amount, t.AccountId)

	if errs != nil {
		logger.Error("Error while creating a new transaction: " + errs.Error())
		tx.Rollback()
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}

	//commit transaction
	errs = tx.Commit()
	if errs != nil {
		logger.Error("Error while committing transaction: " + errs.Error())
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}

	//getting last inserted id
	transactionId, errs := result.LastInsertId()
	if errs != nil {
		logger.Error("Error while getting last inserted id for the new transaction: " + errs.Error())
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}

	//getting the latest account details
	account, appError := d.FindBy(t.AccountId)
	if appError != nil {
		return nil, appError
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	//updating the transaction struct with the lastest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *err.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	errs := d.client.Get(&account, sqlGetAccount, accountId)
	if errs != nil {
		logger.Error("Error while fetching account information: " + errs.Error())
		return nil, err.NewUnexpectedError("Unexpected error from database")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
