package builder

import (
	"builder_design_pattern/builder/house_builder"
)

type iBuilder interface {
	SetWindowType(string)
	SetDoorType(string)
	SetNumFloor(int)
	GetHouse() house_builder.House
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return house_builder.NewNormalBuilder() //&normalBuilder{}
	}

	if builderType == "igloo" {
		return house_builder.NewIglooBuilder() //&iglooBuilder{}
	}
	return nil
}
