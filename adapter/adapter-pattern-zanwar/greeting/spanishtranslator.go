package greeting

import "fmt"

var phrases = map[string]string{"hello": "hola", "how are you?": "comusta?"}

type spanishTranslator struct{}

func NewSpanishTranslator() *spanishTranslator {
	return &spanishTranslator{}
}

func (t *spanishTranslator) translate(phrase string) {
	p, ok := phrases[phrase]
	if !ok {
		fmt.Println("...no spanish translation for:", phrase)
		return
	}
	fmt.Println("Spanish translator:", p)
}
