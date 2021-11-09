package characters

import "fmt"

// BatMan is a concrete type that implements Recite
type BatMan struct{}

func NewBatMan() *BatMan {
	return &BatMan{}
}

// Recite -- BatMan says the dialogue
func (sum BatMan) Recite() {
	fmt.Println("It's not who I am underneath, " +
		"but what I do that defines me!")
}
