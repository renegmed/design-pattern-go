package beverage

import "fmt"

type DarkRoast struct{}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{}
}

func (dr *DarkRoast) Description() string {
	cost := fmt.Sprintf("%.2f", dr.Cost())
	return "Dark Roast Coffee(" + cost + ")"
}

func (dr *DarkRoast) Cost() float32 {
	return 0.99
}
