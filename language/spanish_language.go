package language

func NewSpanishLanguage() *Language {
	return &Language{
		Name: Spanish,
		ToReplace: map[rune]rune{
			'á': 'a',
			'é': 'e',
			'í': 'i',
			'ó': 'o',
			'ú': 'u',
			'ñ': 'N',
		},
	}
}
