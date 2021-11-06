package domain

import (
	"banking/err"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("error while querying customer table" + err.Error())
		return nil, err
	}
	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			log.Println("error while scanning customer table" + err.Error())
			return nil, err
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
			return nil, err.NewNotFoundError("Customer not found")
		} else {
			log.Println("error while scanning customer table" + errs.Error())
			return nil, err.NewUnexpectedError("unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "vulcan:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
