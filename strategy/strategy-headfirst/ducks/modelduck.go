package ducks

import (
	"fmt"
	"strategy-design-pattern/behaviour"
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
