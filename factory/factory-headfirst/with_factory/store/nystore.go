package store

import (
	"factory-design-pattern/pizza"
	"fmt"
)

/**
 * Each subclass implements the createPizza() method
 * and make use of the orderPizza() method in the parent class
 */
type nyPizzaStore struct {
	*aPizzaStore
}

func NewNYPizzaStore() IPizzaStore {
	basePizzaStore := &aPizzaStore{} // abstract (partially implemented interface)

	nyPizzaStore := &nyPizzaStore{basePizzaStore} //assigning the nyPizzaStore's createPizza method to aPizzaStore's createPizza method
	nyPizzaStore.aPizzaStore.CreatePizza = nyPizzaStore.CreatePizza

	return nyPizzaStore
}

func (n *nyPizzaStore) CreatePizza(pizzaType string) (pizza.IPizza, error) {
	switch pizzaType {
	case "cheese":
		return pizza.NewNYStyleCheesePizza(), nil
	case "greek":
		return pizza.NewNYStyleGreekPizza(), nil
	case "pepperoni":
		return pizza.NewNYStylePepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}
