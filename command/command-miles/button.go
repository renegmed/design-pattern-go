package main

type button struct {
    command command
}

func newButton(command command) button {
	return button{command}	 
}

func (b *button) press() {
    b.command.execute()
}