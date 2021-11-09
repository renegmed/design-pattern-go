package genders

import (
	"bridge_design_pattern/bridge"
	"fmt"
)

type female struct {
	name       string
	age        int
	department bridge.Department
}

func NewFemale(name string, age int, department bridge.Department) *female {
	return &female{name, age, department}
}

func (f *female) DescribePerson() string { // interface Gender
	description := fmt.Sprintf("%s is a %d years old woman\n", f.name, f.age)
	description = description + fmt.Sprintf("She works in the %s department", f.department.GetDepartmentName())
	return description
}
