package service

import (
	"banking/domain"
	"banking/dto"
	"banking/err"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *err.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *err.AppError)
	GetAllCustomerByStatus(string) ([]domain.Customer, *err.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *err.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetAllCustomerByStatus(status string) ([]domain.Customer, *err.AppError) {
	return s.repo.FindAllByStatus(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *err.AppError) {
	customer, errs := s.repo.ById(id)

	if errs != nil {
		return nil, errs
	}

	response := dto.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.Status,
	}

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
