package interaction

import (
	"fmt"

	"github.com/hreluz/language-characters-replacer/language"
)

func ShowAvailableActions() {
	fmt.Println("Choose Language: ")
	fmt.Println("-------------------------------")

	for i := 0; i < len(language.LanguageNames); i++ {
		fmt.Printf("(%d) %s\n", (i + 1), language.LanguageNames[i])
	}
}
