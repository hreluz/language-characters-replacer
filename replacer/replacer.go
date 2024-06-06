package replacer

import (
	"strings"

	"github.com/hreluz/language-characters-replacer/language"
)

type Replacer interface {
	Replace(string, language.Language) string
}

func Exec(s string, l language.Language) string {
	var sb strings.Builder

	for _, r := range s {
		if replacement, found := l.ToReplace[r]; found {
			sb.WriteRune(replacement)
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
