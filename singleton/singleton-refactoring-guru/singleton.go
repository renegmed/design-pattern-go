package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single
var once sync.Once

func getInstance() *single {
	once.Do(func() {
		fmt.Println("+++ Creating single instance now.")
		singleInstance = &single{}
	})
	fmt.Println("...Single instance already created.")

	// if singleInstance == nil {
	// 	// lock.Lock()
	// 	// defer lock.Unlock()
	// 	if singleInstance == nil {
	// 		fmt.Println("Creating single instance now.")
	// 		singleInstance = &single{}
	// 	} else {
	// 		fmt.Println("Single instance already created.")
	// 	}
	// } else {
	// 	fmt.Println("Single instance already created.")
	// }

	return singleInstance
}

func (s *single) process() {
	i := rand.Intn(50) + 1
	fmt.Println("Single: processing item no.", i)
}
