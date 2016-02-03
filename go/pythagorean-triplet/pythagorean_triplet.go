package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet
	for t := range iterator(min, max) {
		if isTriplet(t) {
			triplets = append(triplets, t)
		}
	}
	return triplets
}

func Sum(p int) []Triplet {
	var triplets []Triplet
	for a := 1; a <= p/3; a++ {
		for b := a; b <= (p-a)/2; b++ {
			c := p - a - b
			t := Triplet{a, b, c}
			if isTriplet(t) {
				triplets = append(triplets, t)
			}
		}
	}
	return triplets
}

func iterator(min, max int) <-chan Triplet {
	ch := make(chan Triplet)
	go func() {
		for a := min; a <= max; a++ {
			for b := a; b <= max; b++ {
				for c := b; c <= max; c++ {
					ch <- Triplet{a, b, c}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func isTriplet(t Triplet) bool {
	return t[0]*t[0]+t[1]*t[1] == t[2]*t[2]
}
