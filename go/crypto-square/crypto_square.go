package cryptosquare

import (
	"math"
	"strings"
)

const testVersion = 2

func Encode(input string) string {
	s := prepare(input)
	len := squareLen(len(s))
	return encode(s, len)
}

func encode(s string, len int) string {
	ret := make([]string, len)

	for k, v := range s {
		ret[k%len] += string(v)
	}

	return string(strings.Join(ret, " "))
}

func squareLen(len int) int {
	return int(math.Ceil(math.Sqrt(float64(len))))
}

func prepare(input string) string {
	ret := make([]rune, 0, len(input))
	for _, c := range input {
		switch {
		case '0' <= c && c <= '9':
			ret = append(ret, c)
		case 'a' <= c && c <= 'z':
			ret = append(ret, c)
		case 'A' <= c && c <= 'Z':
			ret = append(ret, c-'A'+'a')
		}
	}
	return string(ret)
}
