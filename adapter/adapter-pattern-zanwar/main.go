package main

import (
	"adapter-design-pattern/greeting"
)

type client struct {
	speaker greeting.Speaker
	adapter greeting.Speaker
}

func NewClient(s greeting.Speaker) *client {
	return &client{speaker: s}
}

func (c *client) setAdapter(a greeting.Speaker) {
	c.adapter = a
}

func (c *client) greetInEnglishSpanish(s string) {
	c.speaker.Speak(s)
	c.adapter.Speak(s)
}

func (c *client) greetInEnglishFrench(s string) {
	c.speaker.Speak(s)
	c.adapter.Speak(s)
}

func main() {

	englishSpeaker := greeting.NewEnglishSpeaker()
	client := NewClient(englishSpeaker)
	spanishTranslator := greeting.NewSpanishTranslator()

	spanishAdapter := greeting.NewSpanishSpeakerAdapter(spanishTranslator)

	client.setAdapter(spanishAdapter)
	client.greetInEnglishSpanish("hello")
	client.greetInEnglishSpanish("how are you?")
	client.greetInEnglishSpanish("where is the capitol?")

	frenchTranslator := greeting.NewFrenchTranslator()
	frenchAdapter := greeting.NewFrenchSpeakerAdapter(frenchTranslator)

	client.setAdapter(frenchAdapter)
	client.greetInEnglishFrench("hello")
	client.greetInEnglishFrench("how are you?")
	client.greetInEnglishFrench("where are you going?")
}
