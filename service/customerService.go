package service

import (
	"banking/domain"
	"banking/dto"
	"banking/err"
)

type CustomerService interface {
	GetAllCustomer() ([]*dto.CustomerResponse, *err.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *err.AppError)
	GetAllCustomerByStatus(string) ([]*dto.CustomerResponse, *err.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]*dto.CustomerResponse, *err.AppError) {
	customers, errs := s.repo.FindAll()
	if errs != nil {
		return nil, errs
	}

	response := make([]*dto.CustomerResponse, len(customers))

	for i, c := range customers {
		response[i] = &dto.CustomerResponse{
			Id:          c.Id,
			Name:        c.Name,
			City:        c.City,
			Zipcode:     c.Zipcode,
			DateOfBirth: c.DateOfBirth,
			Status:      c.Status,
		}
	}

	return response, nil
}

func (s DefaultCustomerService) GetAllCustomerByStatus(status string) ([]*dto.CustomerResponse, *err.AppError) {
	customers, errs := s.repo.FindAllByStatus(status)
	if errs != nil {
		return nil, errs
	}

	response := make([]*dto.CustomerResponse, len(customers))

	for i, c := range customers {
		response[i] = &dto.CustomerResponse{
			Id:          c.Id,
			Name:        c.Name,
			City:        c.City,
			Zipcode:     c.Zipcode,
			DateOfBirth: c.DateOfBirth,
			Status:      c.Status,
		}
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *err.AppError) {
	customer, errs := s.repo.ById(id)

	if errs != nil {
		return nil, errs
	}

	response := customer.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
