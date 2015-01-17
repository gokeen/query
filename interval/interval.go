package interval

import (
	"fmt"
)

type Interval string

const (
	Minutely Interval = "minutely"
	Hourly   Interval = "hourly"
	Daily    Interval = "daily"
	Weekly   Interval = "weekly"
	Monthly  Interval = "monthly"
	Yearly   Interval = "yearly"
)

type Reference interface {
	Minute(...int) Interval
	Hour(...int) Interval
	Day(...int) Interval
	Week(...int) Interval
	Month(...int) Interval
	Year(...int) Interval
}

// reference represents the reference any interval derives it's meaning against
// eg. this, previous, next, every, etc...
type reference string

var (
	Every Reference = New("every")
)

// New returns a new interval reference
func New(s string) reference {
	return reference(s)
}

func (t reference) when(f string, n ...int) Interval {
	if len(n) > 0 {
		if i := n[0]; i > 0 {
			return Interval(fmt.Sprintf("%s_%d_%ss", t, i, f))
		}
	}

	return Interval(fmt.Sprintf("%s_%s", t, f))
}

func (t reference) Minute(n ...int) Interval {
	return t.when("minute", n...)
}

func (t reference) Hour(n ...int) Interval {
	return t.when("hour", n...)
}

func (t reference) Day(n ...int) Interval {
	return t.when("day", n...)
}

func (t reference) Week(n ...int) Interval {
	return t.when("week", n...)
}

func (t reference) Month(n ...int) Interval {
	return t.when("month", n...)
}

func (t reference) Year(n ...int) Interval {
	return t.when("year", n...)
}
