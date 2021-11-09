package turkey

type TurkeyAdapter struct {
	turkey turkey
}

func NewTurkeyAdapter(t turkey) *TurkeyAdapter {
	return &TurkeyAdapter{
		turkey: t,
	}
}

func (t *TurkeyAdapter) SetTurkey(turkey turkey) {
	t.turkey = turkey
}

/**
 * Implement the interface of the
 * type you’re adapting to. This is the
 * interface your client expects to see
 */
func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 1; i <= 3; i++ {
		t.turkey.Fly()
	}
}
