package behaviour

import "fmt"

/**
 * Rocket-powered flying behaviour
 */
type FlyRocketPowered struct{}

func (frp *FlyRocketPowered) Fly() {
	fmt.Println("I’m flying with a rocket!")
}
