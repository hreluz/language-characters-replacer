package replacer

import (
	"testing"

	"github.com/hreluz/language-characters-replacer/language"
)

func TestReplace(t *testing.T) {
	lang := language.Language{
		Name: "Invented",
		ToReplace: map[rune]rune{
			'a': 'r',
			'e': 'z',
			'i': 'v',
		},
	}
	var replacer Replacer = ConcreteReplacer{}
	languageFunction := replacer.Replace(lang)
	got := languageFunction("aei")
	want := "rzv"

	if got != want {
		t.Errorf("Error: It was expected %s got %s", want, got)
	}
}
