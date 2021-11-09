package store

import (
	"factory-design-pattern/pizza"
	"fmt"
)

// Abstract Interface
type IPizzaStore interface {
	OrderPizza(pizzaType string) pizza.IPizza
	CreatePizza(pizzaType string) (pizza.IPizza, error)
} // Abstract Concrete Type

type aPizzaStore struct {
	CreatePizza func(pizzaType string) (pizza.IPizza, error)
}

//var _ IPizzaStore = aPizzaStore{}

func (a *aPizzaStore) OrderPizza(pizzaType string) pizza.IPizza {
	var pizz pizza.IPizza
	var err error
	if pizz, err = a.CreatePizza(pizzaType); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	pizz.Prepare()
	pizz.Bake()
	pizz.Cut()
	pizz.Box()
	return pizz
}
