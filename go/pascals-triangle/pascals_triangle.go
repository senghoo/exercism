package pascal

// Triangle pascal traiangle
func Triangle(n int) (res [][]int) {
	for i := 0; i < n; i++ {
		line := make([]int, 0, i+1)
		for j := 0; j <= i; j++ {
			line = append(line, combination(i, j))
		}
		res = append(res, line)
	}
	return
}

func arrangement(n, i int) int {
	return factorial(n) / factorial(n-i)
}

func combination(n, i int) int {
	return arrangement(n, i) / factorial(i)
}

func factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * factorial(x-1)
	}
	return
}
