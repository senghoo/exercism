package leap

const TestVersion = 1

func IsLeapYear(year int) bool {
	yearDivisbleBy := func(x int) bool {
		return (year % x) == 0
	}

	return yearDivisbleBy(4) && (!yearDivisbleBy(100) || yearDivisbleBy(400))
}
