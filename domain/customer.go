package domain

import (
	"banking/dto"
	"banking/err"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	switch c.Status {
	case "1":
		return "active"
	case "0":
		return "inactive"
	default:
		return "unknown"
	}
}

func (c Customer) ToDto() dto.CustomerResponse {
	status := c.statusAsText()
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      status,
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, *err.AppError)
	ById(string) (*Customer, *err.AppError)
	FindAllByStatus(string) ([]Customer, *err.AppError)
}
