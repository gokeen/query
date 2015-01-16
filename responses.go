package query

import (
	"encoding/json"
	"time"
)

// result is the basic response struct for queried results
type result struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

// Result represents results with a numberic base. Most queries will return this
// unless it is a series, extraction or select unique based query
type Result struct {
	result
	Result int64 `json:"result"`
}

type timeframe struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// SeriesResult represents series data from interval based queries
type SeriesResult struct {
	result
	Result []struct {
		Value     int64     `json:"value"`
		Timeframe timeframe `json:"timeframe"`
	} `json:"result"`
}

const (
	KEEN_TIMESTAMP_FMT = "2006-01-02T03:04:05"
)

// parseTs parses the timesstamp to custom format
func parseTs(s string) (time.Time, error) {
	t, err := time.ParseInLocation(KEEN_TIMESTAMP_FMT, s, time.UTC)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

type timeframeRaw struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// UnmarshalJSON overrides to parse keen style timestamps
func (tf *timeframe) UnmarshalJSON(b []byte) error {
	var tr timeframeRaw
	err := json.Unmarshal(b, &tr)
	if err != nil {
		return err
	}

	for _, v := range []struct {
		raw string
		t   *time.Time
	}{
		{tr.Start, &tf.Start},
		{tr.End, &tf.End},
	} {
		t, err := parseTs(v.raw)
		if err != nil {
			return err
		}

		*v.t = t
	}

	return nil
}

// UniqueResult is the basic response struture for the SelectUnique query
type UniqueResult struct {
	result
	Result []interface{} `json:"result"`
}

// ExtractionResult is the basic resposne structure for the Extraction query
// If *email* is defined, *result* will be a string message
type ExtractionResult struct {
	result
	Result interface{} `json:"result"`
}
