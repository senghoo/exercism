package gigasecond

import "time"
import "math"

const TestVersion = 2

func AddGigasecond(t time.Time) time.Time {
	d := time.Duration(math.Pow10(9)) * time.Second
	return t.Add(d)
}

var Birthday = time.Now()
