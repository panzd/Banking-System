package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Jason", "ShangHai", "210001", "2000-01-01", "1"},
		{"1002", "Naci", "ShangHai", "210001", "2000-01-02", "1"},
	}

	return CustomerRepositoryStub{customers}
}
