package luhn

import (
	"strconv"
	"strings"
)

func AddCheck(n string) string {
	sum := check(n + "0")
	if sum == 0 {
		return n + strconv.Itoa(sum)
	}
	return n + strconv.Itoa(10-sum)
}

func Valid(n string) bool {
	return check(n) == 0
}

func check(n string) int {
	p := removeSpace(n)
	len := len(p)
	sum := 0

	//padding
	switch {
	case len == 0:
		return -1
	case len%2 == 1:
		p = "0" + p
	}

	for idx, d := range p {
		if d < '0' || '9' < d {
			return -1
		}

		num := d - '0'
		if idx%2 == 0 {
			num *= 2
			if num >= 10 {
				num = num - 9
			}
		}

		sum += int(num)
	}

	return sum % 10
}

func removeSpace(n string) string {
	return strings.Replace(n, " ", "", -1)
}
