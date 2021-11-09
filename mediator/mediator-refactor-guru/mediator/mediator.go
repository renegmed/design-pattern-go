package mediator

type Mediator interface {
	CanArrive(train) bool
	NotifyAboutDeparture()
}
