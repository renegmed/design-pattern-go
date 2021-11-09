package ducks

import (
	"fmt"
	"strategy-design-pattern/behaviour"
)

type MallardDuck struct {
	*abstractDuck
}

func NewMallardDuck() *MallardDuck { // implementing the abstract method
	d := &abstractDuck{
		Display: func() {
			fmt.Println("I’m a real Mallard duck")
		},
		flyingBehaviour:   &behaviour.FlyWithWings{}, // default flying behaviour
		quackingBehaviour: &behaviour.Quack{},        // default quacking behaviour
	}

	return &MallardDuck{d}
}
