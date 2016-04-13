package lsproduct

import "errors"

const testVersion = 3

func LargestSeriesProduct(input string, span int) (int, error) {
	var max int
	if span < 0 {
		return -1, errors.New("span not valid")
	}
	digits := stringToNumArray(input)
	if digits == nil {
		return -1, errors.New("input not valid")
	}

	digitsLen := len(digits)
	if digitsLen < span {
		return -1, errors.New("span is larger than nums length")
	}
	for i := 0; i <= digitsLen-span; i++ {
		c := multiN(digits[i : i+span]...)
		if c > max {
			max = c
		}
	}
	return max, nil
}

func multiN(nums ...int) int {
	res := 1
	for _, i := range nums {
		res *= i
	}
	return res
}
func stringToNumArray(input string) []int {
	res := make([]int, 0, len(input))
	for _, i := range input {
		if '0' <= i && i <= '9' {
			res = append(res, int(i-'0'))
		} else {
			return nil
		}
	}
	return res
}
