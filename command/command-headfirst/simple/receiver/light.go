package receiver

import "fmt"

type Light struct{}

func (*Light) On() {
	fmt.Println("Light was turned on.")
}

func (*Light) Off() {
	fmt.Println("Light was turned off.")
}
