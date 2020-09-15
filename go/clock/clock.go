package clock

import (
	"strconv"
)

// A Clock handles times without dates
//
// Programs using clocks should typically store and pass them as values,
// not pointers. That is, clock variables and struct fields should be of
// type clock.Clock, not *clock.Clock.
type Clock struct {
	hour   int
	minute int
}

// New creates and returns a Clock
func New(hour, minute int) Clock {
	return Clock{hour: hour, minute: minute}
}

// String returns the clocks time as a string
func (c Clock) String() string {
	c = c.Convert()
	hour := strconv.Itoa(c.hour)
	minute := strconv.Itoa(c.minute)
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}
	return hour + ":" + minute
}

// Add minuets to clock
func (c Clock) Add(minuets int) Clock {
	c.minute = c.minute + minuets
	c.Convert()
	return c
}

// Subtract minuets from clock
func (c Clock) Subtract(minuets int) Clock {
	c.minute = c.minute - minuets
	c.Convert()
	return c
}

// RollOver hour and minute
func (c Clock) RollOver() Clock {
	// roll over hours and minuets
	for c.hour >= 24 || c.minute >= 60 {
		hn := int(c.hour / 24)
		if hn != 0 {
			c.hour = c.hour - (24 * hn)
		}
		mn := int(c.minute / 60)
		if mn != 0 {
			c.hour += mn
			c.minute = c.minute - (60 * mn)
		}
	}
	return c
}

// Convert hours and minuets
func (c Clock) Convert() Clock {
	c = c.RollOver()

	// handle negative  numbers
	for c.hour < 0 || c.minute < 0 {
		if c.hour < 0 {
			c.hour = c.hour + 24
		}
		if c.minute < 0 {
			mn := int(c.minute / 60)
			c.hour += mn
			if mn == 0{
				// can not multiply by 0
				mn = -1
			}
			c.minute = c.minute + (-60 * mn)
		}
		c.RollOver()
	}

	return c
}
