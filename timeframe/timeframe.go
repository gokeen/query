package timeframe

import (
	"gopkg.in/gokeen/query.v1/interval"
	"net/url"
	"time"
)

const (
	isofmt = time.RFC3339
)

const (
	Today     = "today"
	Yesterday = "yesterday"
)

var (
	This     interval.Reference = interval.New("this")
	Previous interval.Reference = interval.New("previous")
)

type absolute struct {
	*url.Values
}

// Absolute returns a new absolute timeframe
func Absolute(opts ...func(a *absolute)) *absolute {
	a := &absolute{
		Values: &url.Values{},
	}
	for _, v := range opts {
		v(a)
	}
	return a
}

func (a absolute) SetStart(t time.Time) {
	a.Set("start", t.Format(isofmt))
}

func (a absolute) SetEnd(t time.Time) {
	a.Set("end", t.Format(isofmt))
}

func (a absolute) Start() (time.Time, error) {
	return time.Parse(isofmt, a.Get("start"))
}

func (a absolute) End() (time.Time, error) {
	return time.Parse(isofmt, a.Get("end"))
}
