package turkey

import "fmt"

type WildTurkey struct{}

func (w *WildTurkey) Gobble() {
	fmt.Println("Wild Turkey: Gobble Gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("Wild Turkey: I'm flying a short distance")
}
