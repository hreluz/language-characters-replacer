package replacer

import (
	"fmt"
	"strings"

	"github.com/hreluz/language-characters-replacer/language"
)

type Replacer interface {
	Replace(string, language.Language) string
}

func Exec(s string, l language.Language) {
	var sb strings.Builder

	for _, r := range s {
		if replacement, found := l.ToReplace[r]; found {
			sb.WriteRune(replacement)
		} else {
			sb.WriteRune(r)
		}
	}

	fmt.Println(sb.String())
}
