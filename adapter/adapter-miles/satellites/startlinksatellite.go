package satellites

import (
	// "adapter-design-pattern/rocket"
	"fmt"
)

type starlinkSatellite struct {
	name string
}

func NewStarlinkSatellite() *starlinkSatellite {
	return &starlinkSatellite{
		name: "Starlink Satellite",
	}
}

func (sls *starlinkSatellite) InsertSatelliteIntoStarlinkPort(f9r *Falcon9Rocket) {
	f9r.payload = append(f9r.payload, sls) // insert myself to payload list as satellite
	fmt.Println("Starlink satellite is attached to Falcon 9 Rocket.")
}
