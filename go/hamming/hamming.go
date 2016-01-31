// +build !example

package hamming

import "errors"

const TestVersion = 2

func Distance(s1, s2 string) (int, error) {
	var diff int
	l := len(s1)
	if l != len(s2) {
		return -1, errors.New("length not match")
	}

	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff, nil
}
