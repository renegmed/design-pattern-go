package satellites

import (
	//"adapter-design-pattern/satellites"
	"fmt"
	// "reflect"
)

type Falcon9Rocket struct {
	payload []interface{}
}

func (f9r *Falcon9Rocket) InsertSatelliteIntoStarlinkPort(s satellite) {
	fmt.Println("Attaching satellite to Falcon 9 Rocket.")
	s.InsertSatelliteIntoStarlinkPort(f9r)
}

func (f9r *Falcon9Rocket) GetPayload() {
	//fmt.Printf("Falcon 9 payload: %v\n", f9r.payload)
	fmt.Println("----- start payload list -----")
	for _, v := range f9r.payload {
		//fmt.Printf("---\n%v\n", v.(*starlinkSatellite).name)
		switch v := v.(type) {
		case *starlinkSatellite:
			//sat := v.(*starlinkSatellite)
			fmt.Println(v.name) //sat.name)
		case *oco2Satellite:
			//sat := v.(*oco2Satellite) //.(oco2Satellite)
			fmt.Println(v.name) //sat.name)
		}
	}
	fmt.Println("----- end payload list -----")

}
