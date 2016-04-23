package strand

const testVersion = 3

func ToRNA(dnaChain string) string {
	res := make([]rune, 0, len(dnaChain))
	for _, dna := range dnaChain {
		switch dna {
		case 'G':
			res = append(res, 'C')
		case 'C':
			res = append(res, 'G')
		case 'T':
			res = append(res, 'A')
		case 'A':
			res = append(res, 'U')
		}
	}
	return string(res)
}
