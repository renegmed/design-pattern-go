package main

import (
	"fmt"
	"log"
)

var pets map[string]Animal

type Animal interface {
	GetName() string
	GetAge() int
	GetSound() string
}

// Pet is a struct that implements Pet interface and
// would be used in any animal struct that we create.
// See `Dog` and `Cat` below
type Pet struct {
	name  string
	age   int
	sound string
}

type Dog struct {
	Pet
}

func (p *Dog) GetName() string {
	return p.name
}

func (p *Dog) GetSound() string {
	return p.sound
}

func (p *Dog) GetAge() int {
	return p.age
}

type Cat struct {
	Pet
}

func (p *Cat) GetName() string {
	return p.name
}

func (p *Cat) GetSound() string {
	return p.sound
}

func (p *Cat) GetAge() int {
	return p.age
}
func GetPet(petType string) (Animal, error) { // returns interface
	if petType == "dog" {
		dog := new(Dog)
		dog.name = "Chester"
		dog.age = 2
		dog.sound = "bark"
		return dog, nil
	}

	if petType == "cat" {
		cat := new(Cat)
		cat.name = "Mr. Buttons"
		cat.age = 3
		cat.sound = "meow"
		return cat, nil
	}

	return nil, fmt.Errorf("Undefined type")
}

func describePet(pet Animal) string {
	return fmt.Sprintf("%s is %d years old. It's sound is %s", pet.GetName(), pet.GetAge(), pet.GetSound())
}

func main() {
	petType := "dog"

	dog, err := GetPet(petType)
	if err != nil {
		log.Fatal(err)
	}
	petDescription := describePet(dog)

	fmt.Println(petDescription)
	fmt.Println("-------------")

	petType = "cat"
	cat, err := GetPet(petType)
	if err != nil {
		log.Fatal(err)
	}
	petDescription = describePet(cat)

	fmt.Println(petDescription)

	fmt.Println("-------------")

	petType = "turtle"
	turtle, err := GetPet(petType)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	petDescription = describePet(turtle)

	fmt.Println(petDescription)
}
