package grains

import "errors"

// Square return how many grains were on each square
func Square(square int) (uint64, error) {
	// calculate 2^(i-1)
	if square > 64 || square < 1 {
		return 0, errors.New("square shound <=64 or >=1")
	}
	i := uint(square)
	return 1 << (i - 1), nil
}

// Total return the total number of grains
func Total() (ret uint64) {
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		ret += s
	}
	return
}
