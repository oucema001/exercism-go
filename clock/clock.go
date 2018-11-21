package clock

import "fmt"

// Clock struct has minutes and hours
type Clock struct {
	hour   int
	minute int
}

//New creates new clock struct with specified minutes and hours
func New(hour int, minute int) Clock {
	m := minute
	if minute < 0 {
		minute = -minute
		m = 60 - (minute % 60)
		hour = hour - (minute+m)/60
	}
	if m >= 60 {

		hour = hour + minute/60
		m = minute % 60
	}
	if hour < 0 {
		hour = 24 + (hour % 24)
	}
	if hour >= 24 {
		hour = hour % 24
	}

	return Clock{
		hour:   hour,
		minute: m,
	}
}

//Add adds specified minutes to clock
func (c Clock) Add(input int) Clock {
	return New(c.hour, c.minute+input)
}

//Subtract removes specified number of minutes from clock
func (c Clock) Subtract(input int) Clock {
	return New(c.hour, c.minute-input)
}

//String formats clock to string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
