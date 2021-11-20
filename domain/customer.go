package domain

import "banking/err"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *err.AppError)
	ById(string) (*Customer, *err.AppError)
	FindAllByStatus(string) ([]Customer, *err.AppError)
}
