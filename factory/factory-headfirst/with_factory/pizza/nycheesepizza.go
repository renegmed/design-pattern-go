package pizza

import "fmt"

type nyStyleCheesePizza struct {
	*pizza
}

func NewNYStyleCheesePizza() IPizza {
	p := &pizza{
		name:     "NY Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}

	return &nyStyleCheesePizza{
		pizza: p,
	}
}

func (p *nyStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into ny style slices")
}
