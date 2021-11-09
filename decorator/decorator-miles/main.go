package main

import "fmt"

func main() {

	// basic pizza plus three toppings
	theBombPizza := &anchovieTopping{
		pizza: &artichokeTopping{
			pizza: &tomatoeTopping{
				pizza: newCheesePizza(),
			},
		},
	}
	
	// basic pizza plus one topping
	pepperoniPizza := &pepperoniTopping{
		pizza: newCheesePizza(),
	}

	fmt.Printf("Price of theBombPizza is $%.2f with %d calories.\n", 
		theBombPizza.getPrice(),
		theBombPizza.getCalories())
	fmt.Printf("Price of pepperoniPizza is $%.2f with %d calories.\n", 
		pepperoniPizza.getPrice(),
		pepperoniPizza.getCalories())
}