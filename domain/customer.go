package domain

import "banking/err"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *err.AppError)
	ById(string) (*Customer, *err.AppError)
}
