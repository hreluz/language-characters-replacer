package replacer

import (
	"strings"

	"github.com/hreluz/language-characters-replacer/language"
)

type ReturnFn func(string) string

type Replacer interface {
	Replace(language.Language) ReturnFn
}

func exec(l language.Language) ReturnFn {
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

// ConcreteReplacer is a concrete type that implements the Replacer interface
type ConcreteReplacer struct{}

// Replace method implements the Replacer interface for ConcreteReplacer
func (cr ConcreteReplacer) Replace(l language.Language) ReturnFn {
	return exec(l)
}
