package helpers

import (
	"strings"
	"unicode"
)

// CleanText applies various cleaning functions to the text.
func CleanText(text string) string {
	text = RemoveDuplicateWords(text)
	text = RemoveDuplicatePunctuation(text)
	return text
}

// RemoveDuplicateWords removes consecutive duplicate words from the text.
func RemoveDuplicateWords(text string) string {
	words := strings.Fields(text)
	if len(words) < 2 {
		return text
	}

	var result []string
	result = append(result, words[0])
	for i := 1; i < len(words); i++ {
		if words[i] != words[i-1] {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}

// RemoveDuplicateCharacters focuses on duplicate punctuation.
func RemoveDuplicatePunctuation(text string) string {
	var result []rune
	for i, r := range text {
		// append if it's the first character or not a punctuation.
		if i == 0 || !unicode.IsPunct(r) {
			result = append(result, r)
			continue
		}

		// For punctuation, check it's not the same as the previous character.
		if unicode.IsPunct(r) && (len(result) == 0 || r != result[len(result)-1]) {
			result = append(result, r)
		}
	}
	return string(result)
}
