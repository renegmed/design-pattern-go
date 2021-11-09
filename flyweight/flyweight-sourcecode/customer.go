package main

//Customer struct
type Customer struct {
	id   string //sequence generator
	name string
	ssn  string
}

func NewCustomer(id string) *Customer {
	return &Customer{id: id}
}

// Customer class method getId
func (customer Customer) getId() string {
	//fmt.Println("getting customer Id")
	return customer.id

}
