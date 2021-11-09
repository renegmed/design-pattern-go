package main

import (
	"fmt"
	"strategy_design_pattern/strategy"
	"strategy_design_pattern/strategy/characters"
)

func main() {

	fmt.Println("+++++++++++++++++++++++++++++++")

	newToy := strategy.NewToy(characters.NewSpiderMan()) //SpiderMan{})
	// This performs the spider man dialogue.
	newToy.PerformDialogue()

	fmt.Println("+++++++++++++++++++++++++++++++")

	// Change the behaviour at runtime.
	newToy.SetSuperHero(characters.NewSuperMan()) //SuperMan{})
	// This performs the super man dialogue.
	newToy.PerformDialogue()

	fmt.Println("+++++++++++++++++++++++++++++++")

	// Change the behaviour at runtime.
	newToy.SetSuperHero(characters.NewBatMan()) //BarMan{})
	// This performs the super man dialogue.
	newToy.PerformDialogue()
}
