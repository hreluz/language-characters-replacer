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
			'ñ': 'n',
			'Á': 'A',
			'É': 'E',
			'Í': 'i',
			'Ó': 'o',
			'Ú': 'u',
			'Ñ': 'N',
		},
	}
}
