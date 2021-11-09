package main

import "fmt"

type doctor struct {
	name string
	next department
}

func newDoctor(name string) *doctor {
	return &doctor{name: name}
}
func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Printf("Doctor %s checkup already done\n", d.name)
		if d.next != nil {
			d.next.execute(p)
		}
		return
	}
	fmt.Printf("Doctor %s checking patient\n", d.name)
	p.doctorCheckUpDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *doctor) setNext(next department) {
	d.next = next
}
