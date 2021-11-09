package main

type EnglishSpeaker struct{}

func (e *EnglishSpeaker) greetInEnglish() string {
	return "Hello there!"
}
