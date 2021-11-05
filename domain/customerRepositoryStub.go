package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Afif", City: "bengaluru", Zipcode: "560085", DateOfBirth: "01-04-2000", Status: "active"},
		{Id: "1002", Name: "Ahmed", City: "Raichur", Zipcode: "584101", DateOfBirth: "24-04-1973", Status: "inactive"},
	}

	return CustomerRepositoryStub{customers}
}
