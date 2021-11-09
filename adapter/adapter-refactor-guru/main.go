package main

import (
	"adapter_design_pattern/adapter"
	"adapter_design_pattern/adapter/mac"
	"adapter_design_pattern/adapter/windows"
)

func main() {

	client := adapter.NewClient()
	mac := mac.NewMac() //&mac.mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := windows.NewWindows()
	windowsMachineAdapter := adapter.NewWindowsAdapter(windowsMachine)

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	dellWindowsMachine := windows.NewDellWindows()
	windowsMachineAdapter = adapter.NewWindowsAdapter(dellWindowsMachine)
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

}
