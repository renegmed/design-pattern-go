package main

import (
	"builder_design_pattern/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	fmt.Println("+++++++++++++++++++++++++")
	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse("Normal window", "Wooden door", 3)

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse("Snow window", "Icy door", 1)

	fmt.Println("+++++++++++++++++++++++++")
	fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}
