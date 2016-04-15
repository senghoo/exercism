package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product int // palindromic, of course
	// list of all possible two-factor factorizations of Product, within
	// given limits, in order
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			num := i * j
			if isPalindrome(num) {
				if pmin.Product == 0 || pmin.Product == num {
					pmin.Product = num
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}

				if pmax.Product == num {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}

				if pmax.Product < num {
					pmax.Product = num
					pmax.Factorizations = [][2]int{[2]int{i, j}}
				}
			}
		}
	}
	if pmin.Product == pmax.Product {
		pmin = Product{}
	}
	if pmax.Product == 0 {
		err = errors.New("No palindromes")
	}
	return
}

func isPalindrome(num int) bool {
	s := strconv.Itoa(num)
	if s == reverseString(s) {
		return true
	}
	return false
}
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
