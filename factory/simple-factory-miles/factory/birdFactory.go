package factory

import (
	"factory_design_pattern/factory/birds"
	"fmt"
)

func MakeBird(birdSpecies string) (IBird, error) {
	switch birdSpecies {
	case "ostrich":
		return birds.NewOstrich(), nil
	case "osprey":
		return birds.NewOsprey(), nil
	default:
		return nil, fmt.Errorf("What kind of bird is that?")
	}
}
