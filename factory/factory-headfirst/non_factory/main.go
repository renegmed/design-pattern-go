package main

import (
	"factory-design-pattern/store"
)

func main() {
	store := store.NewPizzaStore()
	store.OrderPizza("pepperoni")
}
