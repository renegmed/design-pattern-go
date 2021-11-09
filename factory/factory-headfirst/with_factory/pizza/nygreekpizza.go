package pizza

import "fmt"

type nyStyleGreekPizza struct {
	*pizza
}

func NewNYStyleGreekPizza() IPizza {
	p := &pizza{
		name:     "NY Style Sauce and Greek Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Greek"},
	}

	return &nyStyleGreekPizza{
		pizza: p,
	}
}

func (p *nyStyleGreekPizza) Cut() {
	fmt.Println("Cutting the pizza into ny style slices")
}
