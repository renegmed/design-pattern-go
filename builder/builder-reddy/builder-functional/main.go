package main

import "fmt"

// Functional Builder

//Employee struct
type Employee struct {
	name, company, address string
}

type handler func(employee *Employee) error

//EmployeeBuilder struct
type EmployeeBuilder struct {
	actions []handler
}

//Called sets name of the employee
func (b *EmployeeBuilder) Called(value string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) error {
		e.name = value
		return nil
	})
	return b
}

//WorksFor sets company of the employee
func (b *EmployeeBuilder) WorksFor(value string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) error {
		e.company = value
		return nil
	})
	return b
}

//At sets address of the employee
func (b *EmployeeBuilder) At(value string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) error {
		e.address = value
		//return fmt.Errorf("Error at At(%s)", value)
	})
	return b
}

//Build builds the employee object
func (b *EmployeeBuilder) Build() (Employee, error) {
	emp := Employee{}
	for _, a := range b.actions {
		err := a(&emp)
		if err != nil {
			//fmt.Println(err)
			return Employee{}, err
		}
	}
	return emp, nil
}

//NewEmployeeBuilder - constructor
func NewEmployeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{}
}

//RunFunctionalBuilder example
// func RunFunctionalBuilder() {
func main() {
	e := NewEmployeeBuilder()
	employee, err := e.Called("Surya").WorksFor("IBM").At("Bangalore").Build()
	if err != nil {
		fmt.Println("...Error:", err)
		return
	}
	fmt.Println(employee)
}
