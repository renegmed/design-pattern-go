package pizza

import "fmt"

type nyStylePepperoniPizza struct {
	*pizza
}

func NewNYStylePepperoniPizza() IPizza {
	p := &pizza{
		name:     "NY Style Sauce and Pepperoni Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Pepperoni"},
	}

	return &nyStylePepperoniPizza{
		pizza: p,
	}
}

func (p *nyStylePepperoniPizza) Cut() {
	fmt.Println("Cutting the pizza into ny style slices")
}
