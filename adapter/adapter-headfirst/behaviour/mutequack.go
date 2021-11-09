package behaviour

import "fmt"

type MuteQuack struct{}

func (mq *MuteQuack) Quack() {
	fmt.Println("Silence")
}
