package mediator

type train interface {
	Arrive()
	Depart()
	PermitArrival()
}
