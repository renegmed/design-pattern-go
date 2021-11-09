package genders

import (
	"bridge_design_pattern/bridge"
	"fmt"
)

type male struct {
	name       string
	age        int
	department bridge.Department
}

func NewMale(name string, age int, department bridge.Department) *male {
	return &male{name, age, department}
}
func (m *male) DescribePerson() string { // implements interface Gender
	description := fmt.Sprintf("%s is a %d years old man\n", m.name, m.age)
	description = description + fmt.Sprintf("He works in the %s department", m.department.GetDepartmentName())
	return description
}
