package house_builder

type iglooBuilder struct {
	// WindowType string
	// DoorType   string
	// Floor      int
	House
}

func NewIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) SetWindowType(wintype string) {
	b.WindowType = wintype //"Snow Window"
}

func (b *iglooBuilder) SetDoorType(doortype string) {
	b.DoorType = doortype //"Snow Door"
}

func (b *iglooBuilder) SetNumFloor(num int) {
	b.Floor = num
}

func (b *iglooBuilder) GetHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
