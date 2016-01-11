package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	h, m int
}

func Time(hour, minute int) Clock {
	return Clock{hour%24 + minute/60, minute % 60}.normalizing()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	return c.normalizing()
}

func (c Clock) normalizing() Clock {
	c.h += c.m / 60
	c.m %= 60
	if c.m < 0 {
		c.m += 60
		c.h -= 1
	}

	c.h %= 24
	if c.h < 0 {
		c.h += 24
	}
	return c
}
