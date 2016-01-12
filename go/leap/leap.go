package leap

const TestVersion = 1

func IsLeapYear(year int) bool {
	divisbleBy := func(a, b int) bool {
		return (a % b) == 0
	}

	return divisbleBy(year, 4) && (!divisbleBy(year, 100) || divisbleBy(year, 400))
}
