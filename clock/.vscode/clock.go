package main

import "fmt"

type Clock struct {
	hour   int
	minute int
}

//New
func New(hour int, minute int) Clock {
	m := minute

	if m > 60 {

		hour = hour + minute/60
		m = minute % 60
	}
	if minute < 0 {
		minute = -minute
		m = 60 - (minute % 60)

		if minute < 60 {
			hour = hour - 1
		} else {
			hour = hour - (minute+m)/60
		}
	}
	if m == 60 {
		hour++
		m = 0

	}
	h := hour
	if h < 0 {
		h = 24 + (hour % 24)
	}
	if h >= 24 {
		h = hour % 24
	}

	return Clock{
		hour:   h,
		minute: m,
	}
}

func (c Clock) Add(input int) Clock {
	min := c.minute + input
	if min >= 60 {
		h := c.hour + min/60
		m := min % 60
		if h >= 24 {
			return New(h%24, m)
		}
		return New(h, m)
	}
	return New(c.hour, min)
}

func (c Clock) Subtract(input int) Clock {

	return New(c.hour, c.minute-input)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func main() {
	clock := New(10, 30)

	// subtract an hour and a half from it
	clock = clock.Subtract(90)
	fmt.Println(clock.String())
}
