package condiments

import (
	"decorator-design-pattern/beverage"
	"fmt"
)

type Soy struct {
	beverage beverage.Beverage // this is a common interface
}

func NewSoy(b beverage.Beverage) *Soy {
	return &Soy{beverage: b}
}

func (s *Soy) Description() string {
	cost := fmt.Sprintf("%.2f", s.Cost())
	return s.beverage.Description() + ", Soy(" + cost + ")"
}

func (s *Soy) Cost() float32 {
	return s.beverage.Cost() + .15
}
