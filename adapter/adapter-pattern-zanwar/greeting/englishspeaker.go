package greeting

import "fmt"

type englishSpeaker struct{}

func NewEnglishSpeaker() *englishSpeaker {
	return &englishSpeaker{}
}

func (s *englishSpeaker) Speak(str string) {
	fmt.Println("English speaker:", str)
}
