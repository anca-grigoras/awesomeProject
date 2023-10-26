package letters

// VowelsFinder interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// FindVowels MyString implements VowelFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
