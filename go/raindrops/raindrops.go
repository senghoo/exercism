package raindrops

import "strconv"

const TestVersion = 1

type RainDrop struct {
	prime int
	str   string
}

func Convert(input int) string {
	inputDivisbleBy := func(x int) bool {
		return (input % x) == 0
	}

	raindrops := [...]RainDrop{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

	res := ""
	for _, i := range raindrops {
		if inputDivisbleBy(i.prime) {
			res += i.str
		}
	}

	if res == "" {
		res = strconv.Itoa(input)
	}
	return res
}
