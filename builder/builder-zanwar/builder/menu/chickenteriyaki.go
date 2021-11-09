package menu

type chickenTeriyakiBuilder struct {
	Sub
}

func NewChickenTeriyakiBuilder() *chickenTeriyakiBuilder {
	b := chickenTeriyakiBuilder{}
	b.bread = "italian"
	b.hasCheese = true
	b.toppings = []string{"roasted chicken", "olives", "onions", "jalapeños"}
	b.sauces = []string{"chilli", "bbq"}
	return &b
}

// func (c *chickenTeriyakiBuilder) SetBread() {
// 	c.Sub.bread = "italian"
// }

// func (c *chickenTeriyakiBuilder) SetCheese() {
// 	c.Sub.hasCheese = true
// }

// func (c *chickenTeriyakiBuilder) SetToppings() {
// 	c.Sub.toppings = []string{"roasted chicken", "olives", "onions", "jalapeños"}
// }

// func (c *chickenTeriyakiBuilder) SetSauces() {
// 	c.Sub.sauces = []string{"chilli", "bbq"}
// }

// func (c *chickenTeriyakiBuilder) GetSub() Sub {
// 	return c.Sub
// }
