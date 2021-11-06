package service

import (
	"banking/domain"
	"banking/err"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *err.AppError)
	GetCustomer(string) (*domain.Customer, *err.AppError)
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

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *err.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
