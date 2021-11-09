package main

import (
	"fmt"
)

func main() {
	var animal Animal

	animal = NewPig("Porky") // subclass or implementation class
	fmt.Println(animal.Eat("feed"))
	fmt.Println(animal.Eat("feed"))
	fmt.Println(animal.Sleep())
	fmt.Println()

	animal = NewDog("Husky", 2)
	fmt.Println(animal.Eat("Bone"))
	fmt.Println(animal.Sleep())
	fmt.Println(animal.Eat("horse meat"))
	fmt.Println(animal.Eat("dog chow"))
	fmt.Println(animal.Sleep())
}
