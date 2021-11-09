package main

import (
	"bridge_design_pattern/bridge"
	"bridge_design_pattern/departments"
	"bridge_design_pattern/genders"
	"fmt"
)

func main() {
	fmt.Println("-------------")
	// maleDeveloper := genders.Male{
	// 	name:       "John",
	// 	age:        22,
	// 	department: departments.Developer{},
	// }
	maleDeveloper := genders.NewMale("John", 22, departments.Developer{})

	//fmt.Println(maleDeveloper.DescribePerson())
	// fmt.Println("-------------")

	// femalePM := genders.Female{
	// 	name:       "Natalie",
	// 	age:        24,
	// 	department: departments.PM{},
	// }
	femalePM := genders.NewFemale("Natalie", 24, departments.PM{})

	// fmt.Println(femalePM.DescribePerson())

	var list = make(map[string]bridge.Gender, 2)

	list["John"] = maleDeveloper
	list["Natalie"] = femalePM

	for _, emp := range list {
		fmt.Println(emp.DescribePerson())
	}
}
