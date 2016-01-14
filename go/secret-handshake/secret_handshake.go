package secret

var secretCode = []struct {
	code int
	h    string
}{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

var (
	secretLen int = len(secretCode)
	reverse   int = 1 << uint(secretLen)
	allCode   int = reverse - 1
)

func Handshake(val int) []string {
	res := make([]string, 0, 4)

	if val <= 0 || allCode&val == 0 {
		return nil
	}

	for i := 0; i < secretLen; i++ {
		// index in secretCode
		index := i
		// if reverse
		if reverse&val != 0 {
			index = secretLen - i - 1
		}

		if secretCode[index].code&val != 0 {
			res = append(res, secretCode[index].h)
		}
	}
	return res
}
