package main

import (
	"github.com/hreluz/language-characters-replacer/language"
	"github.com/hreluz/language-characters-replacer/replacer"
)

func main() {
	spanishText := "áañeéiíóouú"
	replacer.Exec(spanishText, *language.NewSpanishLanguage())
}
