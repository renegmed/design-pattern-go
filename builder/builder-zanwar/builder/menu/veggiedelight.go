package menu

type veggieDelightBuilder struct {
	Sub
}

func NewVeggieDelightBuilder() *veggieDelightBuilder {
	b := veggieDelightBuilder{}
	b.bread = "parmesan oregano"
	b.hasCheese = false
	b.toppings = []string{"olives", "tomatoes", "onions", "jalapeños"}
	b.sauces = []string{"south west"}
	return &b
}

// func (v *veggieDelightBuilder) SetBread() {
// 	v.Sub.bread = "parmesan oregano"
// }

// func (v *veggieDelightBuilder) SetCheese() {
// 	v.Sub.hasCheese = false
// }

// func (v *veggieDelightBuilder) SetToppings() {
// 	v.Sub.toppings = []string{"olives", "tomatoes", "onions", "jalapeños"}
// }

// func (v *veggieDelightBuilder) SetSauces() {
// 	v.Sub.sauces = []string{"south west"}
// }

// func (v *veggieDelightBuilder) GetSub() Sub {
// 	return v.Sub
// }
