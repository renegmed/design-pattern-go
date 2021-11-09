package main

import "fmt"

// Dog
type Dog struct {
	*animal
}

func NewDog(name string, max int) *Dog {
	p := &animal{
		name: name,
		max:  max,
	}
	return &Dog{p}
}

// Dog's Eat() method, Dog will not continue to eat if he is full
func (d *Dog) Eat(food string) bool {
	eat := func() bool { fmt.Printf("[%v] ate [%v] and continue to eat~\n", d.name, food); return true }
	full := func() bool { fmt.Printf("[%v] has eaten [%v] and is full~\n", d.name, food); return true }
	refuse := func() bool { fmt.Printf("[%v] is full, refuse to eat~\n", d.name); return false }
	return d.animal.eat(food, eat, full, refuse)
}

// Dog's Sleep() method, Dog only goes to sleep when he is full
func (d *Dog) Sleep() bool {
	full := func() bool { fmt.Printf("[%v] is full, go to bed zzz\n", d.name); return true }
	notFull := func() bool { fmt.Printf("[%v] did not eat enough, refused to sleep\n", d.name); return false }
	return d.animal.sleep(full, notFull)
}
