package main

import (
	"builder_design_pattern/builder"
	"fmt"
)

// func describeSub(sub menu.Sub) {
// 	fmt.Printf("bread: %s, cheese: %t, toppings: %s, sauces: %s\n",
// 		sub.GetBread(), sub.HasCheese(), strings.Join(sub.GetToppings(), ", "), strings.Join(sub.GetSauces(), ", "))
// }

func main() {
	veggieDelightBuilder := builder.GetBuilder("veggie delight") //&veggieDelightBuilder{}

	director := builder.NewDirector()
	director.SetBuilder(veggieDelightBuilder)

	// director := &director{
	// 	builder: veggieDelight,
	// }
	veggieDelightSub := director.BuildSub()
	builder.DescribeSub(veggieDelightSub)

	fmt.Println("------------")
	chickenTeriyakiBuilder := builder.GetBuilder("chicken teriyaki")
	director.SetBuilder(chickenTeriyakiBuilder)
	chickenTeriyakiSub := director.BuildSub()

	builder.DescribeSub(chickenTeriyakiSub)

	fmt.Println("------------")
	specialChickenTeriyakiBuilder := builder.GetBuilder("special chicken teriyaki")
	director.SetBuilder(specialChickenTeriyakiBuilder)
	specialChickenTeriyakiSub := director.BuildSub()

	builder.DescribeSub(specialChickenTeriyakiSub)
}
