package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id: "100",
			Name: "One hundred bucks",
			City: "Los Angeles",
			Zipcode: "110011",
			BirthDate: "1997-01-01",
			Status: "1",
		},
		{
			Id: "300",
			Name: "Three hundred bucks",
			City: "New York",
			Zipcode: "220022",
			BirthDate: "1997-02-02",
			Status: "1",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}