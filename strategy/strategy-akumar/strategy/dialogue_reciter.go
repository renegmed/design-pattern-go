package strategy

// DialogueReciter known how to recite a dialogue
type DialogueReciter interface {
	// Concrete types should implement this method.
	Recite()
}
