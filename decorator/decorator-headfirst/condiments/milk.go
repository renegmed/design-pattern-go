package condiments

import (
	"decorator-design-pattern/beverage"
	"fmt"
)

type Milk struct {
	beverage beverage.Beverage // this is a common interface
}

func NewMilk(b beverage.Beverage) *Milk {
	return &Milk{beverage: b}
}
func (m *Milk) Description() string {
	cost := fmt.Sprintf("%.2f", m.Cost())
	return m.beverage.Description() + ", Milk(" + cost + ")"
}

func (m *Milk) Cost() float32 {
	return m.beverage.Cost() + .1
}
