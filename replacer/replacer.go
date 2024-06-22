package replacer

import (
	"strings"

	"github.com/hreluz/language-characters-replacer/language"
)

type ReturnFn func(string) string

type Replacer interface {
	Replace(language.Language) ReturnFn
}

func Exec(l language.Language) ReturnFn {
	return func(s string) string {
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
}
