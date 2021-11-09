package behaviour

import "fmt"

type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("Squeak")
}
