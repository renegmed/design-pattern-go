package main

import "fmt"

func main() {
	// tv := &tv {
	// 	isOn: false,
	// 	volume: 10,
	// }

	// Instantiate commands
	// onCommand := &onCommand{
	// 	device: tv,
	// }
	// offCommand := &offCommand{
	// 	device: tv,
	// }
	// increaseVolumeCommand := &increaseVolumeCommand {
	// 	device: tv,
	// }
	// decreaseVolumeCommand := &decreaseVolumeCommand {
	// 	device: tv,
	// }

	tv := newTv()
	// Instantiate buttons
	onButton := newButton(newOnCommand(tv))
	offButton := newButton(newOffCommand(tv))
	increaseVolumeButton := newButton(newIncreaseVolumeCommand(tv))
	decreaseVolumeButton := newButton(newDecreaseVolumeCommand(tv))

	// onButton := &button{
	// 	command: onCommand,
	// }
	// offButton := &button{
	// 	command: offCommand,
	// }
	// increaseVolumeButton := &button {
	// 	command: increaseVolumeCommand,
	// }
	// decreaseVolumeButton := &button {
	// 	command: decreaseVolumeCommand,
	// }

	// Execute
	increaseVolumeButton.press()
	onButton.press()
	increaseVolumeButton.press()
	decreaseVolumeButton.press()
	offButton.press()
	fmt.Println(tv)
	fmt.Println("++++++++++++++++++++")
	increaseVolumeButton.press()
	onButton.press()
	//fmt.Println(tv)
	fmt.Println("++++++++++++++++++++")
	onButton.press()
	increaseVolumeButton.press()
	//fmt.Println(tv)
	fmt.Println("++++++++++++++++++++")
	increaseVolumeButton.press()
	//fmt.Println(tv)
	fmt.Println("++++++++++++++++++++")
	offButton.press()

}
