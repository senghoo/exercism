package letter

func ConcurrentFrequency(ss []string) FreqMap {
	len := len(ss)
	res := make(FreqMap)
	sem := make(chan FreqMap, len)

	for _, s := range ss {
		go parallelFrequency(s, sem)
	}

	for i := 0; i < len; i++ {
		for k, v := range <-sem {
			res[k] += v
		}
	}
	return res

}

func parallelFrequency(s string, c chan FreqMap) {
	c <- Frequency(s)
}
