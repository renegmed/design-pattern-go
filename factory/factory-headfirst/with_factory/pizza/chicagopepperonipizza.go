package pizza

import "fmt"

type chicagoStylePepperoniPizza struct {
	*pizza
}

func NewChicagoStylePepperoniPizza() IPizza {
	p := &pizza{
		name:     "Chicago Style Deep Dish Pepperoni Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Pepperoni"},
	}

	return &chicagoStylePepperoniPizza{
		pizza: p,
	}
}

/**
 * The Chicago Style Pizza overrides the cut method
 * so that the pieces are cut in squares.
 */
func (c *chicagoStylePepperoniPizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
