package turkey

import (
	"adapter-design-pattern/ducks"
	"fmt"
)

type MallardDuckWrapper struct {
	duck    *ducks.MallardDuck
	adapter *TurkeyAdapter
}

func NewMallardDuckWrapper(d *ducks.MallardDuck, a *TurkeyAdapter) *MallardDuckWrapper {
	return &MallardDuckWrapper{duck: d, adapter: a}
}

func (m *MallardDuckWrapper) Quack() {
	fmt.Println("+++++ duck wrapper quack +++++++")
	fmt.Print("Mallard Duck: ")
	m.duck.PerformQuack()
	m.adapter.Quack() // turkey can now quacks
	fmt.Println("+++++++++++++++++++++++++++++++++")

}

func (m *MallardDuckWrapper) Fly() {
	fmt.Println("+++++ duck wrapper fly +++++++")
	fmt.Print("Mallard Duck: ")
	m.duck.PerformFly()
	m.adapter.Fly()
	fmt.Println("+++++++++++++++++++++++++++++++++")
}
