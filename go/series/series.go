package slice

// All returns a list of all substrings of s with length n.
func All(n int, s string) (res []string) {
	len := len(s)
	for i := 0; i <= len-n; i++ {
		if v, ok := First(n, s[i:]); ok {
			res = append(res, v)
		}
	}
	return
}

// Frist returns the first substring of s with length n.
func Frist(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) || n <= 0 {
		return "", false
	}
	return s[:n], true
}
