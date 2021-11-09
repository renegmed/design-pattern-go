package factory

import (
	"factory_design_pattern/factory/pets"
	"fmt"
)

func MakePet(petType string) (Animal, error) { // returns interface
	if petType == "dog" {
		dog := pets.NewDog("Chester", 2, "bark")
		// dog := new(pets.Dog)
		// dog.name = "Chester"
		// dog.age = 2
		// dog.sound = "bark"
		return dog, nil
	}

	if petType == "cat" {
		cat := pets.NewCat("Mr. Buttons", 3, "meow")
		// cat := new(pets.Cat)
		// cat.name = "Mr. Buttons"
		// cat.age = 3
		// cat.sound = "meow"
		return cat, nil
	}

	return nil, fmt.Errorf("Undefined type")
}
