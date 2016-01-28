package etl

import "strings"

// Transform extract scores from a legacy system.
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for score, letters := range in {
		for _, letter := range letters {
			lower := strings.ToLower(letter)
			out[lower] = score
		}
	}
	return out
}
