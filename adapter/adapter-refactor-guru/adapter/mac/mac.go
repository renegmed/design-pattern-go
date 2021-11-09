package mac

import "fmt"

type mac struct {
}

func NewMac() *mac {
	return &mac{}
}

func (m *mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
