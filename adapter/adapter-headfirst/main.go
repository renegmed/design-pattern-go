package main

import (
	"adapter-design-pattern/ducks"
	"adapter-design-pattern/turkey"
	"fmt"
)

func main() {
	// Creating a duck and a turkey
	mallardDuck := ducks.NewMallardDuck() //MallardDuck{}

	wildTurkey := &turkey.WildTurkey{}

	// Wrapping the turkey in a turkey adapter, which makes it look like a duck
	turkeyAdapter := turkey.NewTurkeyAdapter(wildTurkey)

	mallardDuckWrapper := turkey.NewMallardDuckWrapper(mallardDuck, turkeyAdapter)

	// Passing the duck to a method testDuck() which expects a Duck object
	fmt.Println("The Duck says...")
	testDuck(mallardDuckWrapper)

	// Passing off Turkey as a Duck to the same method
	fmt.Println("\nThe Turkey Adapter says...")
	testDuck(turkeyAdapter)

	fmt.Println("--------------------------------------")

	cityTurkey := &turkey.CityTurkey{}
	turkeyAdapter.SetTurkey(cityTurkey)
	mallardDuckWrapper = turkey.NewMallardDuckWrapper(mallardDuck, turkeyAdapter)
	fmt.Println("The Duck says...")
	testDuck(mallardDuckWrapper)
}

func testDuck(d turkey.Duck) {
	d.Quack() // in case of turkey, it now quacks instead of gobble. Duck as usual, expects to quack
	d.Fly()
}
