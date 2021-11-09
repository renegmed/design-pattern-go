package main

import (
	"factory_design_pattern/factory"
	"fmt"
)

func main() {
	ak47, _ := factory.MakeGun("ak47")
	musket, _ := factory.MakeGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
