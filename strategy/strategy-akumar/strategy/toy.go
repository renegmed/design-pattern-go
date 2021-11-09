package strategy

type toy struct {
	// DialogueReciter is the behaviour that is encapsulated
	// Now this DialogueReciter is of interface type and hence
	// can hold any concrete type.
	// Now the concrete type implements the methods of the
	// DialogueReciter interface.
	DialogueReciter DialogueReciter
}

// NewToy returns a toy object
func NewToy(dr DialogueReciter) *toy {
	return &toy{
		DialogueReciter: dr,
	}
}

// PerformDialogue performs the dialogue
func (t *toy) PerformDialogue() {
	t.DialogueReciter.Recite()
}

// SetSuperHero sets the superhero for the toy
func (t *toy) SetSuperHero(dr DialogueReciter) {
	t.DialogueReciter = dr
}
