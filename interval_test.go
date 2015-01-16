package query

import (
	"gopkg.in/nowk/assert.v2"
	"testing"
)

func TestInterval(t *testing.T) {
	assert.Equal(t, Interval("every_minute"), Every.Minute())
	assert.Equal(t, Interval("every_5_minutes"), Every.Minute(5))

	assert.Equal(t, Interval("every_hour"), Every.Hour())
	assert.Equal(t, Interval("every_5_hours"), Every.Hour(5))

	assert.Equal(t, Interval("every_day"), Every.Day())
	assert.Equal(t, Interval("every_5_days"), Every.Day(5))

	assert.Equal(t, Interval("every_week"), Every.Week())
	assert.Equal(t, Interval("every_5_weeks"), Every.Week(5))

	assert.Equal(t, Interval("every_month"), Every.Month())
	assert.Equal(t, Interval("every_5_months"), Every.Month(5))

	assert.Equal(t, Interval("every_year"), Every.Year())
	assert.Equal(t, Interval("every_5_years"), Every.Year(5))
}

func TestZeroReturnsSingular(t *testing.T) {
	every := NewReference("every")
	assert.Equal(t, Interval("every_month"), every.when("month", 0))
}

func BenchmarkNoInt(b *testing.B) {
	i := 0
	for ; i < b.N; i++ {
		Every.Month()
	}
}

func BenchmarkWithInt(b *testing.B) {
	i := 0
	for ; i < b.N; i++ {
		Every.Month(i)
	}
}

// BechmarkNoInt            2000000               888 ns/op
// BenchmarkWithInt         1000000              1277 ns/op
