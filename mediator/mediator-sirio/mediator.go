package main

type ATCMediator interface {
	Communicate(airTransport *AirTransport, message string)
	Message(airTransport1, airTransport2 *AirTransport, message string)
	Join(at *AirTransport)
}
