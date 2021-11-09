package greeting

import "fmt"

var frenchPhrases = map[string]string{"hello": "bonjour", "how are you?": "comment allez-vous?"}

type frenchTranslator struct{}

func NewFrenchTranslator() *frenchTranslator {
	return &frenchTranslator{}
}

func (t *frenchTranslator) translate(phrase string) {
	p, ok := frenchPhrases[phrase]
	if !ok {
		fmt.Println("...no French translation for:", phrase)
		return
	}
	fmt.Println("french translator:", p)
}
