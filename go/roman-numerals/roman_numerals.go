package romannumerals

import (
	"errors"
	"strings"
)

const testVersion = 2

var (
	numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans  = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
)

func ToRomanNumeral(in int) (res string, err error) {
	if in <= 0 || in >= 4000 {
		err = errors.New("too large or too smail")
		return
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		res += strings.Repeat(romans[i], in/numbers[i])
		in %= numbers[i]
	}
	return
}
