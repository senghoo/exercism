package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	h, m int
}

func Time(hour, minute int) Clock {
	return Clock{hour, minute}.normalizing()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	return Time(c.h, c.m+minutes)
}

func (c Clock) normalizing() Clock {
	timeCarry := func(v, limit int) (int, int) {
		carry := v / limit
		remain := v % limit
		if remain < 0 {
			carry -= 1
			remain += limit
		}
		return remain, carry
	}

	var carry int
	c.m, carry = timeCarry(c.m, 60)
	c.h, _ = timeCarry(c.h+carry, 24)
	return c
}
