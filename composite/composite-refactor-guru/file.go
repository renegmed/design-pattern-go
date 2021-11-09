package main

import "fmt"

type file struct {
	name string
}

func NewFile(name string) *file {
	return &file{name}
}

func (f *file) search(keyword string) { // implement component interface
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}
