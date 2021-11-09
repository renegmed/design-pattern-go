package main

import (
	"mediator_design_pattern/mediator"
	"mediator_design_pattern/mediator/train"
)

func main() {
	stationManager := mediator.NewStationManager()

	// passengerTrain := &passengerTrain{
	// 	mediator: stationManager,
	// }
	// freightTrain := &freightTrain{
	// 	mediator: stationManager,
	// }

	passengerTrain := train.NewPassengerTrain(stationManager)
	freightTrain := train.NewFreightTrain(stationManager)
	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
