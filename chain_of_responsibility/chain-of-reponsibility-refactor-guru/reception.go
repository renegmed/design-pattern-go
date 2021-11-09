package main

import "fmt"

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		if r.next != nil {
			r.next.execute(p)
		}
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *reception) setNext(next department) {
	r.next = next
}
