package main

import "fmt"

func main() {
	game := newGame()

	//Add Terrorist
	err := game.addTerrorist("John", TerroristDressType)
	if err != nil {
		panic(err)
	}

	game.addTerrorist("Peter", TerroristDressType)
	game.addTerrorist("Andrew", TerroristDressType)
	game.addTerrorist("Matt", TerroristDressType)

	//Add CounterTerrorist
	err = game.addCounterTerrorist("Mary", CounterTerrroristDressType)
	if err != nil {
		panic(err)
	}
	game.addCounterTerrorist("Rachel", CounterTerrroristDressType)
	game.addCounterTerrorist("Magda", CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}

	for _, p := range game.getTerrorists() {
		p.describe()
	}

	for _, p := range game.getCounterTerrorists() {
		p.describe()
	}

}
