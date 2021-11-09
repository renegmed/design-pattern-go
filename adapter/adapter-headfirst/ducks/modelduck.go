package ducks

import (
	"adapter-design-pattern/behaviour"
	"fmt"
)

type ModelDuck struct {
	*abstractDuck
}

func NewModelDuck() *ModelDuck { // implementing the abstract method
	d := &abstractDuck{
		Display: func() {
			fmt.Println("I’m a real Model duck")
		},
		flyingBehaviour:   &behaviour.FlyNoWay{},
		quackingBehaviour: &behaviour.MuteQuack{},
	}

	return &ModelDuck{d}
}
