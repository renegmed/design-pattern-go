package main

type mushroomTopping struct {
	pizza pizza
}

func (c *mushroomTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 14
}
