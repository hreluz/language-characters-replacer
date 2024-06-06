package language

type LanguageName string

const (
	Spanish    LanguageName = "Spanish"
	Portuguese LanguageName = "Portuguese"
)

var LanguageNames = [2]LanguageName{
	Spanish,
	Portuguese,
}

type Language struct {
	Name      LanguageName
	ToReplace map[rune]rune
}
