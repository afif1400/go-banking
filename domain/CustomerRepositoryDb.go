package domain

import (
	"banking/err"
	"banking/logger"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *err.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, errs := d.client.Query(findAllSql)

	if errs != nil {
		logger.Error("Error while querying customer table " + errs.Error())
		return nil, err.NewUnexpectedError("unexpected database error")
	}
	var customers []Customer
	for rows.Next() {
		var customer Customer
		errs := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
		if errs != nil {
			logger.Error("Error while scanning customer table " + errs.Error())
			return nil, err.NewNotFoundError("unexpected database error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *err.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers where customer_id = ?"

	row := d.client.QueryRow(findByIdSql, id)
	var customer Customer
	errs := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
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
	rows, errs := d.client.Query(findByStatusSql, status)
	if errs != nil {
		return nil, err.NewUnexpectedError("unexpected database error")
	}
	var customers []Customer
	for rows.Next() {
		var customer Customer
		errs := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
		if errs != nil {
			return nil, err.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	client, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
