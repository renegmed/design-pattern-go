package beverage

import "fmt"

type HouseBlend struct{}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{}
}

func (hb *HouseBlend) Description() string {
	cost := fmt.Sprintf("%.2f", hb.Cost())
	return "House Blend Coffee(" + cost + ")"
}

func (hb *HouseBlend) Cost() float32 {
	return 0.89
}
