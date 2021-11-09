package airplane

import (
	"fmt"
	"mediator_design_pattern/mediator"
)

type airbusA380 struct {
	mediator mediator.Mediator
}

func NewAirbusA380(mediator mediator.Mediator) *airbusA380 {
	return &airbusA380{mediator}
}
func (g *airbusA380) RequestArrival() {
	if g.mediator.CanLand(g) {
		fmt.Println("AirbusA380: Landing")
	} else {
		fmt.Println("AirbusA380: Waiting")
	}
}

func (g *airbusA380) Departure() {
	fmt.Println("AirbusA380: Leaving")
	g.mediator.NotifyFree()
}

func (g *airbusA380) PermitArrival() {
	fmt.Println("AirbusA380: Arrival Permitted. Landing")
}
