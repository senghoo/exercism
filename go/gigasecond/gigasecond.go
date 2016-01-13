package gigasecond

import "time"

const TestVersion = 2

func AddGigasecond(t time.Time) time.Time {
	d := time.Duration(1e9) * time.Second
	return t.Add(d)
}

var Birthday = time.Now()
