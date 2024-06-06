package language

type LanguageName string

const (
	Spanish LanguageName = "Spanish"
)

type Language struct {
	Name      LanguageName
	ToReplace map[rune]rune
}
