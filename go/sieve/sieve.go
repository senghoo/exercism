package sieve

func Sieve(upto int) (res []int) {
	nums := twoToN(upto)

	for n := 2; n <= upto; n++ {
		if !nums[n] {
			continue
		}
		for i := 1; i*n <= upto; i++ {

			nums[i*n] = false
		}
		res = append(res, n)
	}
	return
}

func twoToN(n int) (res map[int]bool) {
	res = make(map[int]bool)
	for i := 2; i <= n; i++ {
		res[i] = true
	}
	return
}
