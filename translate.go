package igpay

import "strings"

const (
	pigLatinSuffix string = "ay"
	vowels         string = "aeiou"
	vowelsSuffix   string = "d" + pigLatinSuffix
)

func isCapitalized(word string) bool {
	first := word[0:1]
	return !(strings.ToLower(first) == first)
}

func capitalize(word string) string {
	return strings.ToUpper(word[0:1]) + word[1:]
}

func startsWithVowel(word string) bool {
	first := strings.ToLower(word[0:1])
	return strings.Contains(vowels, first)
}

func suffix(word string) (suffix string) {
	if startsWithVowel(word) {
		suffix = vowelsSuffix
	} else {
		suffix = strings.ToLower(word[0:1]) + pigLatinSuffix
	}
	return
}

func base(word string) (base string) {
	if startsWithVowel(word) {
		base = word[0:1] + strings.ToLower(word[1:])
	} else {
		base = strings.ToLower(word[1:])
		if isCapitalized(word) {
			base = capitalize(base)
		}
	}
	return
}

func Translate(in string) (out string) {
	inWords := strings.Fields(in)
	outWords := make([]string, 0)

	for _, inWord := range inWords {
		outWord := base(inWord) + suffix(inWord)
		outWords = append(outWords, outWord)
	}

	out = strings.Join(outWords, " ")
	return
}
