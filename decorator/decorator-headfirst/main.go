package main

import (
	"decorator-design-pattern/beverage"
	"decorator-design-pattern/condiments"
	"fmt"
)

func main() {
	fmt.Println("-------------------------------")
	/**
	 * Order up an espresso, no condiments
	 * and print its description and cost
	 */
	espresso := &beverage.Espresso{} //&espresso{}

	fmt.Printf("%s $%.2f\n", espresso.Description(), espresso.Cost())

	fmt.Println("-------------------------------")

	// Make a DarkRoast object.
	darkRoast := &beverage.DarkRoast{} // &darkRoast{}

	// Wrap it with a Mocha
	// singleMocha := &condiments.Mocha{darkRoast}
	singleMocha := condiments.NewMocha(darkRoast)

	// Wrap it in a second Mocha
	doubleMocha := condiments.NewMocha(singleMocha)
	// doubleMocha := &condiments.Mocha{
	// 	beverage: singleMocha,
	// }

	// Wrap it in a Whip.
	doubleMochaWhip := condiments.NewWhip(doubleMocha)

	// doubleMochaWhip := &condiments.Whip{
	// 	beverage: doubleMocha,
	// }

	fmt.Printf("%s $%.2f\n", doubleMochaWhip.Description(), doubleMochaWhip.Cost())

	fmt.Println("-------------------------------")

	// Finally give us a HouseBlend with Soy, Mocha, and Whip.

	houseBlend := beverage.NewHouseBlend()
	soy := condiments.NewSoy(houseBlend)
	mocha := condiments.NewMocha(soy)
	soyMochaWhipHouseBlend := condiments.NewWhip(mocha)

	// soyMochaWhipHouseBlend := &condiments.Whip{
	// 	beverage: &condiments.Mocha{
	// 		beverage: &condiments.Soy{
	// 			beverage: &beverage.HouseBlend{},
	// 		},
	// 	},
	// }

	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.Description(), soyMochaWhipHouseBlend.Cost())
	fmt.Println("-------------------------------")
}
