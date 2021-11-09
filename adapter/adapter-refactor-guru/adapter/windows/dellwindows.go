package windows

import "fmt"

type dellWindows struct{}

func NewDellWindows() *dellWindows {
	return &dellWindows{}
}
func (w *dellWindows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into special Dell windows machine.")
}
