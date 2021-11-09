package main

//Employee struct
type Employee struct {
	id   string
	name string
}

func NewEmployee(id string) *Employee {
	return &Employee{id: id}
}

//Employee class method getId
func (employee Employee) getId() string {
	return employee.id
}
