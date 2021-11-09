package main

import (
	"fmt"
	"strategy-design-pattern/behaviour"
	"strategy-design-pattern/ducks"
)

func main() {
	mallardDuck := ducks.NewMallardDuck()

	mallardDuck.Display()
	mallardDuck.PerformFly()
	mallardDuck.PerformQuack()
	mallardDuck.Swim()

	modelDuck := ducks.NewModelDuck()
	modelDuck.Display()
	modelDuck.PerformFly()
	modelDuck.PerformQuack()
	modelDuck.Swim()

	// --- new (altered) flying behaviour of model duck ----

	fmt.Println("++++++++++++++++++++")

	modelDuck.SetFlyingBehaviour(&behaviour.FlyRocketPowered{})
	fmt.Print("New Flying Behaviour of Model Duck: ")
	modelDuck.PerformFly()

	mallardDuck.SetFlyingBehaviour(&behaviour.FlyRocketPowered{})
	fmt.Print("New Flying Behaviour of Mallard Duck: ")
	mallardDuck.PerformFly()
	mallardDuck.SetQuackingBehaviour(&behaviour.MuteQuack{})
	fmt.Print("New Quacking Behaviour of Mallard Duck: ")
	modelDuck.PerformQuack()
}
