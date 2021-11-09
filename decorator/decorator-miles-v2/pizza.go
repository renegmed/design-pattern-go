package main

type pizza interface {
	getPrice() float64
	getCalories() int
}

type cheesePizza struct {
	calories int
	price    float64
}

func newCheesePizza(price float64, calories int) *cheesePizza {
	return &cheesePizza{
		calories: calories,
		price:    price,
	}
}

func (c *cheesePizza) getPrice() float64 {
	return c.price
}

func (c *cheesePizza) getCalories() int {
	return c.calories
}
