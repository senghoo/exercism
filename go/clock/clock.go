package clock

import "fmt"

const TestVersion = 2

type Clock int

func NewClock(minutes int) Clock {
	minutes_in_a_day := 60 * 24
	minutes %= minutes_in_a_day
	if minutes < 0 {
		minutes += minutes_in_a_day
	}
	return Clock(minutes)
}

func Time(hour, minute int) Clock {
	return NewClock(hour*60 + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	return NewClock(int(c) + minutes)
}
