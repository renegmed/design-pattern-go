package factory

type IBird interface {
	SetName(name string)
	GetName() string
	TryToFly()
}
