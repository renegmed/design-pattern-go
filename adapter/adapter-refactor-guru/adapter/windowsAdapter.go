package adapter

import (
	"adapter_design_pattern/adapter/windows"
	"fmt"
)

type windowsAdapter struct {
	//windowMachine *windows.Windows
	windowMachine windows.IWindows
}

func NewWindowsAdapter(windows windows.IWindows) *windowsAdapter {
	return &windowsAdapter{windows}
}
func (w *windowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.InsertIntoUSBPort()
}
