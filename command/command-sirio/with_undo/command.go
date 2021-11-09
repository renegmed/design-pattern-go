package main

type Command interface {
	Call()
	Undo()
}
