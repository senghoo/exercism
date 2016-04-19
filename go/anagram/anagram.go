package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (res []string) {
	lSubject := strings.ToLower(subject)
	s := StringToOrderdRuneList(lSubject)
	for _, c := range candidates {
		lc := strings.ToLower(c)
		if Eq(s, StringToOrderdRuneList(lc)) && lc != lSubject {
			res = append(res, lc)
		}
	}
	return
}

func StringToOrderdRuneList(input string) []int {
	runes := make([]int, 0, len(input))
	for _, r := range input {
		runes = append(runes, int(r))
	}
	sort.Ints(runes)
	return runes
}

func Eq(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
