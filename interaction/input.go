package interaction

import (
	"fmt"
	"strconv"

	"github.com/hreluz/language-characters-replacer/language"
)

func GetActionChoice() language.LanguageName {

	for {
		choice, _ := GetUserInput("Your action choice: ")

		c, _ := strconv.Atoi(choice)

		if c <= 0 || c > len(language.LanguageNames) {
			fmt.Println("Incorrect choice, try again")
			continue
		}

		return language.LanguageNames[c-1]
	}
}
