package receiver

import "fmt"

type Light struct {
	Name string
}

func (l *Light) On() {
	fmt.Println(l.Name, "light was turned on.")
}

func (l *Light) Off() {
	fmt.Println(l.Name, "light was turned off.")
}
