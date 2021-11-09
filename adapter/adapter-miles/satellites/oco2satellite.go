package satellites

import (
	"fmt"
)

type oco2Satellite struct {
	name string
}

func NewOco2Satellite() *oco2Satellite {
	return &oco2Satellite{
		name: "OCO2 Satellite",
	}
}

func (oco2s *oco2Satellite) insertSatelliteIntoOco2Port(f9r *Falcon9Rocket) {
	f9r.payload = append(f9r.payload, oco2s) // insert myself to payload list as satellite
	fmt.Println("OCO2 satellite is attached to Falcon 9 Rocket.")
}
