package builder

import "builder_design_pattern/builder/house_builder"

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) BuildHouse(wintype, doortype string, floor int) house_builder.House {
	d.builder.SetDoorType(doortype)
	d.builder.SetWindowType(wintype)
	d.builder.SetNumFloor(floor)
	return d.builder.GetHouse()
}
