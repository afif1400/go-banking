package domain

import (
	"banking/err"
	"banking/logger"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *err.AppError) {
	var customers []Customer
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	errs := d.client.Select(&customers, findAllSql)
	// rows, errs := d.client.Query(findAllSql)

	if errs != nil {
		logger.Error("Error while querying customer table " + errs.Error())
		return nil, err.NewUnexpectedError("unexpected database error")
	}

	// errs = sqlx.StructScan(rows, &customers)
	// if errs != nil {
	// 	logger.Error("Error while scanning customer table " + errs.Error())
	// 	return nil, err.NewNotFoundError("unexpected database error")
	// }

	//? this is the naive way of scanning using the default package
	// for rows.Next() {
	// 	var customer Customer
	// 	errs := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
	// 	if errs != nil {
	// 		logger.Error("Error while scanning customer table " + errs.Error())
	// 		return nil, err.NewNotFoundError("unexpected database error")
	// 	}
	// 	customers = append(customers, customer)
	// }

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *err.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers where customer_id = ?"

	var customer Customer
	errs := d.client.Get(&customer, findByIdSql, id)
	if errs != nil {
		if errs == sql.ErrNoRows {
			logger.Error("Customer not found")
			return nil, err.NewNotFoundError("Customer not found")
		} else {
			logger.Error("error while scanning customer table" + errs.Error())
			return nil, err.NewUnexpectedError("unexpected database error")
		}
	}

	return &customer, nil
}

func (d CustomerRepositoryDb) FindAllByStatus(status string) ([]Customer, *err.AppError) {
	findByStatusSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers where status = ? "
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		return nil, err.NewNotFoundError("Not customer with status " + status)
	}
	var customers []Customer
	errs := d.client.Select(&customers, findByStatusSql, status)
	if errs != nil {
		return nil, err.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "vulcan:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
