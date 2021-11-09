package train

import (
	"fmt"
	"mediator_design_pattern/mediator"
)

type freightTrain struct {
	mediator mediator.Mediator
}

func NewFreightTrain(mediator mediator.Mediator) *freightTrain {
	return &freightTrain{mediator}
}

func (g *freightTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *freightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.NotifyAboutDeparture()
}

func (g *freightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.Arrive()
}
