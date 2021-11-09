package receiver

import (
	"fmt"
	"strconv"
)

type Stereo struct {
	Name string
}

func (s *Stereo) On() {
	fmt.Println(s.Name, "stereo is on.")
}

func (s *Stereo) Off() {
	fmt.Println(s.Name, "stereo is off.")
}

func (s *Stereo) SetCD() {
	fmt.Println(s.Name, "CD is set.")
}

func (s *Stereo) SetVolume(v int) {
	fmt.Println(s.Name, "Volume is set to", strconv.Itoa(v))
}
