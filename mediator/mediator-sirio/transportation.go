package main

import "fmt"

type AirTransport struct {
	PilotsName string
	Type       string
	Runway     ATCMediator
	runwayLog  []string
}

func NewHelicopter(pilotsName string) *AirTransport {
	return &AirTransport{PilotsName: pilotsName, Type: "Helicopter"}
}

func NewPlane(pilotsName string) *AirTransport {
	return &AirTransport{PilotsName: pilotsName, Type: "Plane"}
}

func NewUFO(pilotsName string) *AirTransport {
	return &AirTransport{PilotsName: pilotsName, Type: "UFO"}
}

func (at *AirTransport) Receive(sender *AirTransport, message string) {
	s := fmt.Sprintf("%s (Transport: %s): '%s'", sender.PilotsName, sender.Type, message)
	fmt.Printf("[%s's runway log] %s\n", sender.PilotsName, s)
	at.runwayLog = append(at.runwayLog, s)
}

func (at *AirTransport) Say(message string) {
	at.Runway.Communicate(at, message)
}

func (at *AirTransport) PrivateMessage(dst *AirTransport, message string) {
	if at != dst {
		at.Runway.Message(at, dst, message)
	}
}
