package main

type EnglishToSpanishAdapter struct {
	speaker EnglishSpeaker
}

func (a *EnglishToSpanishAdapter) greetInSpanish() string {
	englishMessage := a.speaker.greetInEnglish()
	return translate(englishMessage)
}

func translate(engMessage string) (spanishMessage string) {
	// *insert complex translation logic*
	spanishMessage = "¡Hola!"
	return
}
