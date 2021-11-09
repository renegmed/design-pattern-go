package main

import (
	"composite-pattern/composite"
	"composite-pattern/employees"
	"composite-pattern/items"
	"fmt"
)

func main() {
	// showEmployeeHierarchy()
	showItemCategoryHierarchy()
}

func showEmployeeHierarchy() {
	ceo := employees.NewEmployee("ID-1", "Socrates")
	vpIdealist := employees.NewEmployee("ID-2", "Plato")
	vpRealist := employees.NewEmployee("ID-3", "Aristotle")
	directorIdealist := employees.NewEmployee("ID-4", "Hegel")
	directorRealist := employees.NewEmployee("ID-5", "Hume")

	directorIdealistOffice := composite.NewComposite(directorIdealist)
	directorRealistOffice := composite.NewComposite(directorRealist)

	vpIdealistOffice := composite.NewComposite(vpIdealist)
	vpIdealistOffice.Add(directorIdealistOffice)

	vpRealistOffice := composite.NewComposite(vpRealist)
	vpRealistOffice.Add(directorRealistOffice)

	company := composite.NewComposite(ceo)
	company.Add(vpRealistOffice)
	company.Add(vpIdealistOffice)

	fmt.Println("=====Display Office of VP of Idealist=====")
	vpIdealistOffice.Display()
	fmt.Println("=====Display Office of VP of Realist=====")
	vpRealistOffice.Display()
	fmt.Println("=====Display Company=====")
	company.Display()
}

func showItemCategoryHierarchy() {
	womenRedShoe := items.NewShoppingItem("Red Shoe", "Red shoe for women", 90)
	womenBlackShoe := items.NewShoppingItem("Black Shoe", "Black shoe for women", 90)

	menRedShoe := items.NewShoppingItem("Red Shoe", "Red shoe for men", 80)
	menBlackShoe := items.NewShoppingItem("Black Shoe", "Black shoe for men", 80)

	womenShoeCategory := composite.NewComposite(items.NewItemCategory("Shoe", "Shoe for women"))
	womenShoeCategory.Add(womenRedShoe)
	womenShoeCategory.Add(womenBlackShoe)

	menShoeCategory := composite.NewComposite(items.NewItemCategory("Shoe", "Shoe for men"))
	menShoeCategory.Add(menRedShoe)
	menShoeCategory.Add(menBlackShoe)

	fmt.Println("===Shoe for Ladies===")
	womenShoeCategory.Display()
	fmt.Println("===Shoe for Gents===")
	menShoeCategory.Display()
}
