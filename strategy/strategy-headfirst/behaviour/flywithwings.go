package behaviour

import "fmt"

/**
 * Flying behaviour implementation for ducks that fly
 */
type FlyWithWings struct{}

func (fw *FlyWithWings) Fly() {
	fmt.Println("I'm flying")
}
