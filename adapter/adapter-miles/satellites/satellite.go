package satellites

// import (
// 	"adapter-design-pattern/rocket"
// )

type satellite interface {
	InsertSatelliteIntoStarlinkPort(f9r *Falcon9Rocket)
}
