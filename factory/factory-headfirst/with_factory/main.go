package main

import (
	"factory-design-pattern/store"
	"fmt"
)

func main() {
	nyPizzaStore := store.NewNYPizzaStore()
	pizza := nyPizzaStore.OrderPizza("cheese")

	fmt.Printf("Ethan ordered %s pizza\n\n", pizza.GetName())

	chicagoPizzaStore := store.NewChicagoPizzaStore()
	pizza = chicagoPizzaStore.OrderPizza("cheese")

	fmt.Printf("Joel ordered %s pizza\n\n", pizza.GetName())
}
