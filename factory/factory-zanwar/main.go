package main

import (
	"factory_design_pattern/factory"
	"fmt"
	"log"
)

var pets map[string]factory.Animal

func describePet(pet factory.Animal) string {
	return fmt.Sprintf("%s is %d years old. It's sound is %s", pet.GetName(), pet.GetAge(), pet.GetSound())
}

func main() {

	fmt.Println("-------------")

	petType := "dog"

	dog, err := factory.MakePet(petType)
	if err != nil {
		log.Fatal(err)
	}
	petDescription := describePet(dog)

	fmt.Println(petDescription)
	fmt.Println("-------------")

	petType = "cat"
	cat, err := factory.MakePet(petType)
	if err != nil {
		log.Fatal(err)
	}
	petDescription = describePet(cat)

	fmt.Println(petDescription)

	fmt.Println("------ this cause an error -------")

	petType = "turtle"
	turtle, err := factory.MakePet(petType)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	petDescription = describePet(turtle)

	fmt.Println(petDescription)
}
