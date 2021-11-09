package items

import "fmt"

type ItemCategory struct {
	name        string
	description string
}

func (i ItemCategory) Display() {
	fmt.Printf("name %s; desc: %s\n", i.name, i.description)
}

func NewItemCategory(name, desc string) ItemCategory {
	return ItemCategory{
		name:        name,
		description: desc,
	}
}
