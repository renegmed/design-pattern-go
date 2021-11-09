package airplane

import (
	"fmt"
	"mediator_design_pattern/mediator"
)

type boeing737 struct {
	mediator mediator.Mediator
}

func NewBoeing737(mediator mediator.Mediator) *boeing737 {
	return &boeing737{mediator}
}

func (g *boeing737) RequestArrival() {
	if g.mediator.CanLand(g) {
		fmt.Println("Boeing737: Landing")
	} else {
		fmt.Println("Boeing737: Waiting")
	}
}

func (g *boeing737) Departure() {
	fmt.Println("Boeing737: Leaving")
	g.mediator.NotifyFree()
}

func (g *boeing737) PermitArrival() {
	fmt.Println("Boeing737: Arrival Permitted. Landing")
}
