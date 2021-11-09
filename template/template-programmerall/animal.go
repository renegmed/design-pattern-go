package main

// abstract animal
type animal struct {
	name string
	food []string
	max  int
	full bool
}

// Eat
// eat-the specific method for not eating enough, full-the specific method
//     for becoming full after eating, refuse-the specific method for refusing to eat
// return whether you have eaten
// There are only combinations in Go. In order to realize the empty method for
// specific Pig/Dog in Animal, the method of passing in func is adopted here.
func (a *animal) eat(food string, eat, full, refuse func() bool) bool {
	if !a.full {
		a.food = append(a.food, food)
		// not full yet
		if len(a.food) < a.max {
			return eat() // leave it to specific animals to implement
		} else {
			a.full = true
			return full() // leave it to specific animals to achieve
		}
	}

	return refuse() // leave it to specific animals to achieve
}

// go to bed
// full-the corresponding operation when full, notFull-the corresponding
//      operation when not full
func (a *animal) sleep(full, notFull func() bool) bool {
	if a.full {
		return full() // leave it to specific animals to achieve
	}
	return notFull() // leave it to specific animals to implement

}
