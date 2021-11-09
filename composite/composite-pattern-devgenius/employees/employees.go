package employees

import "fmt"

type Employee struct {
	id   string
	name string
}

func (e Employee) Display() {
	fmt.Printf("ID: %s\nName: %s\n", e.id, e.name)
}

func NewEmployee(id, name string) Employee {
	return Employee{
		id:   id,
		name: name,
	}
}
