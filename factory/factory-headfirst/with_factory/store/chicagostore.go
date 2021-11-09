package store

import (
	"factory-design-pattern/pizza"
	"fmt"
)

type chicagoPizzaStore struct {
	*aPizzaStore
}

// var _ IPizzaStore = NewChicagoPizzaStore() //chicagoPizzaStore{}

func NewChicagoPizzaStore() IPizzaStore {
	basePizzaStore := &aPizzaStore{} // abstract pizza store

	chicagoPizzaStore := &chicagoPizzaStore{basePizzaStore}

	chicagoPizzaStore.aPizzaStore.CreatePizza = chicagoPizzaStore.CreatePizza // assign the CreatePizza function to struct field

	return chicagoPizzaStore
}

// this method complete the implementation of IPizzaStore interface
func (c *chicagoPizzaStore) CreatePizza(pizzaType string) (pizza.IPizza, error) {
	switch pizzaType {
	case "cheese":
		return pizza.NewChicagoStyleCheesePizza(), nil
	case "greek":
		return pizza.NewChicagoStyleGreekPizza(), nil
	case "pepperoni":
		return pizza.NewChicagoStylePepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}
