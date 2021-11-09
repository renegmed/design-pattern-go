package behaviour

import "fmt"

/**
 * Flying behaviour implementation for ducks
 * that do not fly (like rubber and decoy ducks)
 */
type FlyNoWay struct{}

func (fnw *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
