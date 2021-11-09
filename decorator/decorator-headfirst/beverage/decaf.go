package beverage

import "fmt"

type Decaf struct{}

func NewDecaf() *Decaf {
	return &Decaf{}
}

func (d *Decaf) Description() string {
	cost := fmt.Sprintf("%.2f", d.Cost())
	return "Decaf(" + cost + ")"
}

func (d *Decaf) Cost() float32 {
	return 1.05
}
