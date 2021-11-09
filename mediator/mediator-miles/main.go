package main

import (
	"mediator_design_pattern/mediator"
	"mediator_design_pattern/mediator/airplane"
)

func main() {
	airTower := mediator.NewAirTower()

	boeing737 := airplane.NewBoeing737(airTower)
	airbus380 := airplane.NewAirbusA380(airTower)

	boeing737.RequestArrival()
	airbus380.RequestArrival()
	boeing737.Departure()
	airbus380.Departure()
}
