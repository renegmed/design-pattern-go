package receiver

import (
	"fmt"
	"strconv"
)

type Stereo struct{}

func (s *Stereo) On() {
	fmt.Println("Stereo is on.")
}

func (s *Stereo) Off() {
	fmt.Println("Stereo is off.")
}

func (s *Stereo) SetCD() {
	fmt.Println("CD is set.")
}

func (s *Stereo) SetVolume(v int) {
	fmt.Println("Volume is set to", strconv.Itoa(v))
}
