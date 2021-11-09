package pizza

import "fmt"

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	PizzaType string
}

func (p *Pizza) Prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.PizzaType)
}

func (p *Pizza) Bake() {
	fmt.Printf("Baking %s Pizza\n", p.PizzaType)
}

func (p *Pizza) Cut() {
	fmt.Printf("Cutting %s Pizza\n", p.PizzaType)
}

func (p *Pizza) Box() {
	fmt.Printf("Boxing %s Pizza\n", p.PizzaType)
}

type CheesePizza struct {
	*Pizza
}

func NewCheesePizza() IPizza {
	p := &Pizza{
		PizzaType: "Cheese",
	}

	return &CheesePizza{
		Pizza: p,
	}
}

// so goes the implementation for greek and pepperoni pizzas

type GreekPizza struct {
	*Pizza
}

func NewGreekPizza() IPizza {
	p := &Pizza{
		PizzaType: "Greek",
	}

	return &GreekPizza{
		Pizza: p,
	}
}

type PepperoniPizza struct {
	*Pizza
}

func NewPepperoniPizza() IPizza {
	p := &Pizza{
		PizzaType: "Pepperoni",
	}

	return &PepperoniPizza{
		Pizza: p,
	}
}
