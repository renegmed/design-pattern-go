package condiments

import (
	"decorator-design-pattern/beverage"
	"fmt"
)

type Whip struct {
	beverage beverage.Beverage // this is a common interface
}

func NewWhip(b beverage.Beverage) *Whip {
	return &Whip{beverage: b}
}

func (w *Whip) Description() string {
	cost := fmt.Sprintf("%.2f", w.Cost())
	return w.beverage.Description() + ", Whip(" + cost + ")"
}

func (w *Whip) Cost() float32 {
	return w.beverage.Cost() + .1
}
