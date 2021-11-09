package customer

import "fmt"

//Customer struct
type Customer struct {
	name string
	id   int
}

func NewCustomer() *Customer {
	return &Customer{}
}

// //Customer class method create - create Customer given nam
func (customer *Customer) Create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

func (customer *Customer) GetName() string {
	return customer.name
}
