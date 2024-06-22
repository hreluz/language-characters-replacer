package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hreluz/language-characters-replacer/interaction"
	"github.com/hreluz/language-characters-replacer/language"
	"github.com/hreluz/language-characters-replacer/replacer"
)

func main() {
	option := ""

	languageFn, err := getLanguageFn()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	for option != "N" {
		option = execute(languageFn)
	}
}

func getLanguageFn() (replacer.ReturnFn, error) {
	interaction.ShowLanguagesActions()
	chosenLanguage := interaction.GetActionChoice()
	lF := func(s string) string { return "" }

	switch chosenLanguage {
	case language.Spanish:
		lF = replacer.Exec(*language.NewSpanishLanguage())
	default:
		errorString := fmt.Sprintf("%v is not implemented.\n", chosenLanguage)
		return lF, errors.New(errorString)
	}

	return lF, nil
}

func execute(lFn replacer.ReturnFn) string {
	originalString, _ := interaction.GetUserInput("Enter your string: ")
	modifiedString := lFn(originalString)

	fmt.Printf("Original string: %s\n", originalString)
	fmt.Printf("Replaced string: %s \n", modifiedString)

	option, _ := interaction.GetUserInput("Would you like to transform another string (Y/N): ")
	return option
}
