package main

import (
	"command-design-pattern/command"
	"command-design-pattern/receiver"
	"fmt"
)

func main() {
	remoteControl := command.NewRemoteControl()

	// Instantiate all the devices
	livingRoomLight := &receiver.Light{"Living Room"}
	kitchenLight := &receiver.Light{"Kitchen"}
	ceilingFan := &receiver.CeilingFan{"Living Room"}
	garage := &receiver.Garage{}
	stereo := &receiver.Stereo{"Living Room"}

	// Instantiate all the Command Objects
	livingRoomLightOn := command.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := command.NewLightOffCommand(livingRoomLight)
	kitchenLightOn := command.NewLightOnCommand(kitchenLight)
	kitchenLightOff := command.NewLightOffCommand(kitchenLight)

	ceilingFanOn := command.NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := command.NewCeilingFanOffCommand(ceilingFan)

	stereoOnWithCD := command.NewStereoOnCommand(stereo)
	stereoOff := command.NewStereoOffCommand(stereo)

	garageDoorOpen := command.NewGarageDoorOpenCommand(garage)
	garageDoorClose := command.NewGarageDoorCloseCommand(garage)

	// Load commands into the remote slots
	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.SetCommand(3, stereoOnWithCD, stereoOff)
	remoteControl.SetCommand(4, garageDoorOpen, garageDoorClose)

	fmt.Println(remoteControl)

	//  We step through each slot and push its On and Off button.
	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
	remoteControl.OnButtonWasPushed(4)
	remoteControl.OffButtonWasPushed(4)
}
