package main

type SpeakSpanish interface {
	greetInSpanish() string
}

// Notice that the SpanishSpeaker already implements the
// SpeakSpanish interface
type SpanishSpeaker struct{}

func (s *SpanishSpeaker) greetInSpanish() string {
	return "¡Hola!"
}
