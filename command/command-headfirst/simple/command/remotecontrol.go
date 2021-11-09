package command

type simpleRemoteControl struct {
	slot ICommand
}

func NewSimpleRemoteControl(c ICommand) *simpleRemoteControl {
	return &simpleRemoteControl{c}
}

/**
 * We have a method for setting
 * the command the slot is going
 * to control. This could be called
 * multiple times if the client of
 * this code wanted to change the
 * behaviour of the remote button.
 */
func (r *simpleRemoteControl) SetCommand(command ICommand) {
	r.slot = command
}

/**
 * This method is called when the
 * button is pressed. All we do is take
 * the current command bound to the
 * slot and call its execute() method.
 */
func (r *simpleRemoteControl) ButtonWasPressed() {
	r.slot.Execute()
}
