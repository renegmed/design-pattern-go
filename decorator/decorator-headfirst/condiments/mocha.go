package condiments

import (
	"decorator-design-pattern/beverage"
	"fmt"
)

type Mocha struct {
	beverage beverage.Beverage // this is a common interface
}

func NewMocha(b beverage.Beverage) *Mocha {
	return &Mocha{beverage: b}
}
func (m *Mocha) Description() string {
	cost := fmt.Sprintf("%.2f", m.Cost())
	return m.beverage.Description() + ", Mocha(" + cost + ")"
}

func (m *Mocha) Cost() float32 {
	return m.beverage.Cost() + .2
}
