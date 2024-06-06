package main

import (
	"fmt"
	"os"

	"github.com/hreluz/language-characters-replacer/interaction"
	"github.com/hreluz/language-characters-replacer/language"
	"github.com/hreluz/language-characters-replacer/replacer"
)

func main() {
	originalString, _ := interaction.GetUserInput("Enter your string: ")
	var modifiedString string

	interaction.ShowAvailableActions()
	choice := interaction.GetActionChoice()

	switch choice {
	case language.Spanish:
		modifiedString = replacer.Exec(originalString, *language.NewSpanishLanguage())
	default:
		fmt.Printf("%v is not implemented.\n", choice)
		os.Exit(1)
	}

	fmt.Printf("Original string: %s ", originalString)
	fmt.Printf("\nReplaced string: %s ", modifiedString)
}
