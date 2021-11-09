package windows

import "fmt"

type Windows struct{}

func NewWindows() *Windows {
	return &Windows{}
}
func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
