package characters

import "fmt"

// SuperMan is a concrete type that implements Recite
type SuperMan struct{}

func NewSuperMan() *SuperMan {
	return &SuperMan{}
}

// Recite -- SuperMan says the dialogue
func (sum SuperMan) Recite() {
	fmt.Println("There is a superhero in all of us, " +
		"we just need the courage to put on the cape")
}
