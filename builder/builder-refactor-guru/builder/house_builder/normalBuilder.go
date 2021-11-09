package house_builder

type normalBuilder struct {
	// WindowType string
	// DoorType   string
	// Floor      int
	House
}

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) SetWindowType(wintype string) {
	b.WindowType = wintype //"Wooden Window"
}

func (b *normalBuilder) SetDoorType(doortype string) {
	b.DoorType = doortype // "Wooden Door"
}

func (b *normalBuilder) SetNumFloor(num int) {
	b.Floor = num
}

func (b *normalBuilder) GetHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
