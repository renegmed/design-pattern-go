package main

//Address struct
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func NewAddress(id string) *Address {
	return &Address{id: id}
}

//Address class method getId
func (address Address) getId() string {
	return address.id
}
