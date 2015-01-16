package query

import (
	"encoding/json"
	"gopkg.in/nowk/assert.v2"
	"strings"
	"testing"
	"time"
)

func TestEmbed(t *testing.T) {
	b := `{
		"error_code": "error",
		"message": "you got keened",
		"result": 123
	}`

	r := strings.NewReader(b)

	var d Result
	err := json.NewDecoder(r).Decode(&d)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "error", d.ErrorCode)
	assert.Equal(t, "you got keened", d.Message)
	assert.Equal(t, int64(123), d.Result)
}

func TestTimeframeUnmarshal(t *testing.T) {
	b := `{"result": [
		{
			"value": 123,
			"timeframe": {
				"start": "2012-08-05T00:00:00",
				"end": "2012-08-06T00:00:00"
			}
		}
	]}`

	r := strings.NewReader(b)

	var s SeriesResult
	err := json.NewDecoder(r).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	tf := s.Result[0].Timeframe

	assert.Equal(t, time.Date(2012, 8, 5, 0, 0, 0, 0, time.UTC), tf.Start)
	assert.Equal(t, time.Date(2012, 8, 6, 0, 0, 0, 0, time.UTC), tf.End)
}
