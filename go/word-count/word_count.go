package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {
	res := make(Frequency)
	splitFunc := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return false
		}
		return true
	}
	for _, word := range strings.FieldsFunc(phrase, splitFunc) {
		res[strings.ToLower(word)]++
	}
	return res
}
