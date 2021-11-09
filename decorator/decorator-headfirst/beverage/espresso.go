package beverage

import "fmt"

type Espresso struct{}

func NewEspresso() *Espresso {
	return &Espresso{}
}

func (e *Espresso) Description() string {
	cost := fmt.Sprintf("%.2f", e.Cost())
	return "Espresso(" + cost + ")"
}

func (e *Espresso) Cost() float32 {
	return 1.99
}
