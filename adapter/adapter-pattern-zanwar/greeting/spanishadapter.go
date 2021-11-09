package greeting

type spanishSpeakerAdapter struct {
	translator translator
}

func NewSpanishSpeakerAdapter(t translator) *spanishSpeakerAdapter {
	return &spanishSpeakerAdapter{t}
}

func (a *spanishSpeakerAdapter) Speak(s string) {
	a.translator.translate(s)
}
