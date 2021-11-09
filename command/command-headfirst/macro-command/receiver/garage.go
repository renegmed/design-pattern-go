package receiver

import "fmt"

type Garage struct {
	Name string
}

func (*Garage) Up() {
	fmt.Println("Garage door is up.")
}

func (*Garage) Down() {
	fmt.Println("Garage door is down.")
}

func (*Garage) Stop() {
	fmt.Println("Garage door stopped moving.")
}

func (*Garage) LightOn() {
	fmt.Println("Garage light was turned on.")
}

func (*Garage) LightOff() {
	fmt.Println("Garage light was turned off.")
}
