package main

type Runway struct {
	airTransports []*AirTransport
}

func NewRunway() *Runway {
	return &Runway{}
}

func (r *Runway) Communicate(src *AirTransport, message string) {
	for _, at := range r.airTransports {
		if at != src {
			at.Receive(at, message)
		}
	}
}

func (r *Runway) Message(src, dst *AirTransport, message string) {
	for _, at := range r.airTransports {
		if at == dst {
			at.Receive(src, message)
		}
	}
}

func (r *Runway) Join(at *AirTransport) {
	joinMsg := at.PilotsName + " joined the Runway"
	r.Communicate(at, joinMsg)

	at.Runway = r
	r.airTransports = append(r.airTransports, at)
}
