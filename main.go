package main

import (
	"fmt"

	"github.com/hreluz/language-characters-replacer/interaction"
	"github.com/hreluz/language-characters-replacer/language"
	"github.com/hreluz/language-characters-replacer/replacer"
)

func main() {
	originalString, _ := interaction.GetUserInput("Enter your string: ")
	modifiedString := replacer.Exec(originalString, *language.NewSpanishLanguage())

	fmt.Printf("Original string: %s ", originalString)
	fmt.Printf("\nReplaced string: %s ", modifiedString)
}
