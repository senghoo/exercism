package raindrops

import "strconv"

const TestVersion = 1

type rainDrop struct {
	prime int
	str   string
}

func Convert(input int) string {
	DivisbleBy := func(x int) bool {
		return (input % x) == 0
	}

	raindrops := [...]rainDrop{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

	var res string
	for _, i := range raindrops {
		if DivisbleBy(i.prime) {
			res += i.str
		}
	}

	if res == "" {
		res = strconv.Itoa(input)
	}
	return res
}
