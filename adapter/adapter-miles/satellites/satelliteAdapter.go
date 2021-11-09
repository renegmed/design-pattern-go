package satellites

import (
	// "adapter-design-pattern/satellites"
	"fmt"
)

type oco2SatelliteAdapter struct {
	oco2Satellite *oco2Satellite
}

func NewOco2SatelliteAdapter(o *oco2Satellite) *oco2SatelliteAdapter {
	return &oco2SatelliteAdapter{o}
}

func (oco2sa *oco2SatelliteAdapter) InsertSatelliteIntoStarlinkPort(f9r *Falcon9Rocket) {
	fmt.Println("Satellite adapter converts Starlink port to OCO2 port.")
	oco2sa.oco2Satellite.insertSatelliteIntoOco2Port(f9r) // add the satellite to payload list kept my falcon9Roocket
}
