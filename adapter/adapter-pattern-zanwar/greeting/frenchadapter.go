package greeting

type frenchSpeakerAdapter struct {
	translator translator
}

func NewFrenchSpeakerAdapter(t translator) *frenchSpeakerAdapter {
	return &frenchSpeakerAdapter{t}
}

func (a *frenchSpeakerAdapter) Speak(s string) {
	a.translator.translate(s)
}
