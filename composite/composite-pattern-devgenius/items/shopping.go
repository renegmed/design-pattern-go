package items

import "fmt"

type ShoppingItem struct {
	name        string
	description string
	price       float64
}

func (s ShoppingItem) Display() {
	fmt.Printf("name: %s; desc: %s; price: %.2f\n", s.name, s.description, s.price)
}

func NewShoppingItem(name, desc string, price float64) ShoppingItem {
	return ShoppingItem{
		name:        name,
		description: desc,
		price:       price,
	}
}
