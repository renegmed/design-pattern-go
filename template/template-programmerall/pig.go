package main

import "fmt"

// Pig
type Pig struct {
	*animal
}

func NewPig(name string) *Pig {
	p := &animal{
		name: name,
	}
	return &Pig{p}
}

// Pig’s Eat() method, Pig will continue to eat whether he is full or not
func (p *Pig) Eat(food string) bool {
	eat := func() bool { fmt.Printf("[%v] ate [%v] and continue to eat~\n", p.name, food); return true }
	return p.animal.eat(food, eat, eat, eat)
}

// Pig’s Sleep() method, Pig can go to sleep whether he is full or not
func (p *Pig) Sleep() bool {
	full := func() bool { fmt.Printf("[%v]Go to bed zzz\n", p.name); return true }
	return p.animal.sleep(full, full)
}
