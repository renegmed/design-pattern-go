package builder

import (
	"builder_design_pattern/builder/menu"
	"fmt"
	"strings"
)

type iSubBuilder interface {
	SetBread(string)
	SetCheese(bool)
	SetToppings([]string)
	SetSauces([]string)
	AddTopping(string)
	AddSauce(string)
	GetSub() menu.Sub
}

func GetBuilder(builderType string) iSubBuilder {
	if builderType == "veggie delight" {
		return menu.NewVeggieDelightBuilder()
	}

	if builderType == "chicken teriyaki" {
		return menu.NewChickenTeriyakiBuilder()
	}

	if builderType == "special chicken teriyaki" {
		b := menu.NewChickenTeriyakiBuilder()
		b.SetBread("Wheat bread")
		return b
	}
	return nil
}

func DescribeSub(sub menu.Sub) {
	fmt.Printf("bread: %s, cheese: %t, toppings: %s, sauces: %s\n",
		sub.GetBread(), sub.HasCheese(), strings.Join(sub.GetToppings(), ", "), strings.Join(sub.GetSauces(), ", "))
}
