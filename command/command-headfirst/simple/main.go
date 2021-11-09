package main

import (
	"command-design-pattern/command"
	"command-design-pattern/receiver"
	"fmt"
)

func main() {

	/**
	 * Now we create a Light
	 * object, this will be the
	 * Receiver of the request
	 */
	light := &receiver.Light{}

	/**
	 * Create a command and
	 * pass the Receiver to it.
	 */
	lightOnCommand := command.NewLightOnCommand(light)

	/**
	 * The remote is our Invoker;
	 * it will be passed a
	 * command object that can
	 * be used to make requests.
	 */
	remoteControl := command.NewSimpleRemoteControl(lightOnCommand)

	/**
	 * Pass the command to the Invoker
	 */
	// remoteControl.setCommand(lightOnCommand)

	/**
	 * We simulate the button being pressed.
	 */
	remoteControl.ButtonWasPressed()

	fmt.Println("++++++++++++++++++")

	garage := &receiver.Garage{}

	garageDoorOpenCommand := command.NewGarageDoorOpenCommand(garage)

	/**
	 * Pass the new command to the invoker
	 */
	remoteControl.SetCommand(garageDoorOpenCommand)

	remoteControl.ButtonWasPressed()
}
