package pizza

import "fmt"

/**
 * We will start with the abstract implement of Pizza.
 * All concrete Pizzas will derive from this.
 */
type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

// Implementing the default preparation steps for pizza: prepare, bake, cut, box
func (p *pizza) Prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}

func (p *pizza) Bake() {
	fmt.Println("Baking Pizza for 25 min at 350")
}

// func (p *pizza) Cut() {
// 	fmt.Println("Cutting the pizza into diagonal slices")
// }

func (p *pizza) Box() {
	fmt.Println("“Place pizza in official PizzaStore box")
}
func (p *pizza) GetName() string {
	return p.name
}
