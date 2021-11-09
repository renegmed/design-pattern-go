package characters

import "fmt"

// SpiderMan is a concrete type that implements Recite
type SpiderMan struct{}

func NewSpiderMan() *SpiderMan {
	return &SpiderMan{}
}

// Recite -- SpiderMan says the dialogue
func (spm SpiderMan) Recite() {
	fmt.Println("No Man Can Win Every Battle, " +
		"But No Man Should Fall Without A Struggle")
}
