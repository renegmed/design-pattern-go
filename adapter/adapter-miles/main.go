package main

import (
	// "adapter-design-pattern/rocket"
	"adapter-design-pattern/satellites"
)

func main() {

	// falcon9Rocket is our client, it expects a Starlink satellite
	falcon9Rocket := &satellites.Falcon9Rocket{}
	falcon9Rocket.GetPayload() // this just prints out the current listed satellites

	// instantiate satellites
	starlinkSatellite := satellites.NewStarlinkSatellite()

	// insert the expected Starlink satellite
	falcon9Rocket.InsertSatelliteIntoStarlinkPort(starlinkSatellite) // message: Attaching satellite to Falcon 9 Rocket. & Starlink satellite is attached to Falcon 9 Rocket.
	falcon9Rocket.GetPayload()                                       // message: Starlink Satellite

	oco2Satellite := satellites.NewOco2Satellite()

	// allow a falcon9Rocket to accept an oco2Satellite via adapter
	oco2SatelliteAdapter := satellites.NewOco2SatelliteAdapter(oco2Satellite)

	// oco2SatelliteAdapter := &satellites.Oco2SatelliteAdapter{
	// 	oco2Satellite: oco2Satellite,
	// }

	// use the adapter to insert an OCO2 satellite. Falcon9Rocket deals with the adapter instead with Oco2Satellite
	falcon9Rocket.InsertSatelliteIntoStarlinkPort(oco2SatelliteAdapter)
	falcon9Rocket.GetPayload() // message: Starlink Satellite & OCO2 Satellite meaning falcon9Rocket can now work with both satellites

}
