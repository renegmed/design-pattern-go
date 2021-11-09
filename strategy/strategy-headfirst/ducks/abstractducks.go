package ducks

import (
	"fmt"
	"strategy-design-pattern/behaviour"
)

type abstractDuck struct {
	Display func()
	/**
	 * Declare two reference variables for the behaviour interface types.
	 * All duck subclasses inherit these
	 */
	flyingBehaviour   behaviour.FlyBehaviour
	quackingBehaviour behaviour.QuackBehaviour
}

/**
 * Delegate fly and quack behaviour to
 * respective behaviour classes
 */
func (d *abstractDuck) PerformFly() {
	d.flyingBehaviour.Fly()
}

func (d *abstractDuck) PerformQuack() {
	d.quackingBehaviour.Quack()
}

func (d *abstractDuck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}

/**
 * Setting the behaviour dynamically
 */
func (d *abstractDuck) SetFlyingBehaviour(fb behaviour.FlyBehaviour) {
	d.flyingBehaviour = fb
}

func (d *abstractDuck) SetQuackingBehaviour(qb behaviour.QuackBehaviour) {
	d.quackingBehaviour = qb
}
