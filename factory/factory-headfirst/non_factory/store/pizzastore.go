package store

import (
	"factory-design-pattern/pizza"
)

type pizzaStore struct{}

func NewPizzaStore() *pizzaStore {
	return &pizzaStore{}
}

/**
 * We are passing the type of pizza
 */
func (ps *pizzaStore) OrderPizza(pizzaType string) {
	var pizz pizza.IPizza

	/**
	 * Based on the type of pizza,
	 * we instantiate the correct concrete class
	 * and assign it to the pizza instance variable.
	 * Note that each pizza instance here has to implement the iPizza interface
	 */
	switch pizzaType {
	case "cheese":
		pizz = pizza.NewCheesePizza()
	case "greek":
		pizz = pizza.NewGreekPizza()
	case "pepperoni":
		pizz = pizza.NewPepperoniPizza()
	}

	pizz.Prepare()
	pizz.Bake()
	pizz.Cut()
	pizz.Box()
}
