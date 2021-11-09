package mediator

type Mediator interface {
	CanLand(airplane) bool
	NotifyFree()
}
