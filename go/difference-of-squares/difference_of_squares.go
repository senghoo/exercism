package diffsquares

func SquareOfSums(count int) int {
	// sum of 1 to count
	sum := (1 + count) * count / 2
	return sum * sum
}

func SumOfSquares(count int) int {
	// sum of squares
	return count * (count + 1) * (2*count + 1) / 6
}

func Difference(count int) int {
	return SquareOfSums(count) - SumOfSquares(count)
}
