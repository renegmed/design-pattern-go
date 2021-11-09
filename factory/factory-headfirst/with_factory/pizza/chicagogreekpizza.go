package pizza

import "fmt"

type chicagoStyleGreekPizza struct {
	*pizza
}

func NewChicagoStyleGreekPizza() IPizza {
	p := &pizza{
		name:     "Chicago Style Deep Dish Greek Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Greek"},
	}

	return &chicagoStyleGreekPizza{
		pizza: p,
	}
}

/**
 * The Chicago Style Pizza overrides the cut method
 * so that the pieces are cut in squares.
 */
func (c *chicagoStyleGreekPizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
