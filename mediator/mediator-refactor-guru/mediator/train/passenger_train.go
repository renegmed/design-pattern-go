package train

import (
	"fmt"
	"mediator_design_pattern/mediator"
)

type passengerTrain struct {
	mediator mediator.Mediator
}

func NewPassengerTrain(mediator mediator.Mediator) *passengerTrain {
	return &passengerTrain{mediator}
}

func (g *passengerTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *passengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.NotifyAboutDeparture()
}

func (g *passengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}
