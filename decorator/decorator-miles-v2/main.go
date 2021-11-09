package main

import "fmt"

func main() {

	var prices = make(map[string]float64)
	prices["cheesepizza"] = 10.00
	prices["pepperoni"] = 1.75
	prices["artichoke"] = 2.00
	prices["anchovie"] = 2.50
	prices["tomato"] = 0.75

	var calories = make(map[string]int)
	calories["cheesepizza"] = 500
	calories["pepperoni"] = 200
	calories["artichoke"] = 150
	calories["anchovie"] = 350
	calories["tomato"] = 50

	var pizzas = make(map[string]pizza)
	theBombPizza := &anchovieTopping{
		toppingprice: prices["anchovie"],
		calories:     calories["anchovie"],
		pizza: &artichokeTopping{
			toppingprice: prices["artichoke"],
			calories:     calories["artichoke"],
			pizza: &tomatoeTopping{
				toppingprice: prices["tomato"],
				calories:     calories["tomato"],
				pizza:        newCheesePizza(prices["cheesepizza"], calories["cheesepizza"]),
			},
		},
	}

	pizzas["bombpizza"] = theBombPizza

	pepperoniPizza := &pepperoniTopping{
		pizza: newCheesePizza(prices["cheesepizza"], calories["cheesepizza"]),
	}
	pizzas["pepperoni"] = pepperoniPizza

	printPizza := func(name string, pz map[string]pizza) {
		p, ok := pz[name]
		if !ok {
			fmt.Printf("Pizza %s not on the record.\n", name)
			return
		}
		fmt.Printf("Price of %s is $%.2f with %d calories.\n",
			name,
			p.getPrice(),
			p.getCalories())
	}

	printPizza("bombpizza", pizzas)
	printPizza("pepperoni", pizzas)
	printPizza("xxxxxxxxx", pizzas)

}
