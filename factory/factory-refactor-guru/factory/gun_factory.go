package factory

import (
	"factory_design_pattern/factory/guns"
	"fmt"
)

func MakeGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return guns.NewAk47(), nil
	}
	if gunType == "musket" {
		return guns.NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
