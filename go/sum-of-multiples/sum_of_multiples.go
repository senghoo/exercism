package summultiples

// MultipleSummer retrun sum of all the multiples of particular numbers up to but not including that number.
func MultipleSummer(v ...int) func(int) int {
	divisible := func(num int) bool {
		for _, v := range v {
			if num%v == 0 {
				return true
			}
		}
		return false
	}

	return func(limit int) (sum int) {
		for n := 1; n < limit; n++ {
			if divisible(n) {
				sum += n
			}
		}
		return
	}
}
