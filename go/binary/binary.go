package binary

import "errors"

func ParseBinary(in string) (res int, err error) {
	for _, r := range in {
		res <<= 1
		switch r {
		case '1':
			res++
		case '0':
		default:
			return 0, errors.New("unexpected char")
		}
	}
	return
}
