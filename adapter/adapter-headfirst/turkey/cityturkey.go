package turkey

import "fmt"

type CityTurkey struct{}

func (w *CityTurkey) Gobble() {
	fmt.Println("City Turkey: Let my phone do the gobbling")
}

func (w *CityTurkey) Fly() {
	fmt.Println("City Turkey: I'm too fat to fly 10 meters")
}
