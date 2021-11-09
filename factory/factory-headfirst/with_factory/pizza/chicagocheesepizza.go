package pizza

import "fmt"

type chicagoStyleCheesePizza struct {
	*pizza
}

func NewChicagoStyleCheesePizza() IPizza {
	p := &pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}

	return &chicagoStyleCheesePizza{
		pizza: p,
	}
}

/**
 * The Chicago Style Pizza overrides the cut method
 * so that the pieces are cut in squares.
 */
func (c *chicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
