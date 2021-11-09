package main

import "fmt"

func main() {

	pizza := &veggeMania{}

	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}
	//Add tomato topping
	pizzaWithCheeseAndTomatoAndMushrrom := &mushroomTopping{
		pizza: pizzaWithCheeseAndTomato,
	}

	fmt.Printf("Price of veggeMania with tomato, cheese and mushroom topping is %d\n",
		pizzaWithCheeseAndTomatoAndMushrrom.getPrice())
}
