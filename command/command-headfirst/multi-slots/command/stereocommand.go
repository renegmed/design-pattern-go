package command

import (
	"command-design-pattern/receiver"
)

/**
 * Just like the LightOnCommand, we get
 * passed the instance of the stereo we
 * are going to be controlling and we store
 * it in a local instance variable.
 */
type stereoOnCommand struct {
	Stereo *receiver.Stereo
}

func NewStereoOnCommand(s *receiver.Stereo) *stereoOnCommand {
	return &stereoOnCommand{
		Stereo: s,
	}
}

/**
 * To carry out this request, we need to call three
 * methods on the stereo: first, turn it on, then set
 * it to play the CD, and finally set the volume to 11.
 */
func (s *stereoOnCommand) Execute() {
	s.Stereo.On()
	s.Stereo.SetCD()
	s.Stereo.SetVolume(11)
}

// -------------------
type stereoOffCommand struct {
	Stereo *receiver.Stereo
}

func NewStereoOffCommand(s *receiver.Stereo) *stereoOffCommand {
	return &stereoOffCommand{
		Stereo: s,
	}
}

/**
 * To carry out this request, we need to call two
 * methods on the stereo: first, turn it off,
 * and set the volume to 0.
 */
func (s *stereoOffCommand) Execute() {
	s.Stereo.Off()
	s.Stereo.SetVolume(0)
}
