package mediator

type airplane interface {
	RequestArrival()
	Departure()
	PermitArrival()
}
