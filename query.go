package query

type queryOpts func(*Query)

// Query is the struct around the properties of keen.io's query resources
type Query struct {
	EventCollection string      `json:"event_collection"`
	TargetProperty  string      `json:"target_property,omitempty"`
	Percentile      float64     `json:"percentile,omitempty"`
	Filters         []*Filter   `json:"filters,omitempty"`
	Timeframe       interface{} `json:"timeframe,omitempty"`
	Interval        Interval    `json:"interval,omitempty"`
	Timezone        int         `json:"timezone,omitempty"`
	GroupBy         string      `json:"group_by,omitempty"`

	Email         string   `json:"email,omitempty"`
	Latest        int      `json:"latest,omitempty"`
	PropertyNames []string `json:"property_names,omitempty"`

	queryType string
	projectID string
}

func new(pid, qtype, coll, targ string, opts ...queryOpts) *Query {
	q := &Query{
		EventCollection: coll,
		TargetProperty:  targ,

		queryType: qtype,
		projectID: pid,
	}

	for _, v := range opts {
		v(q)
	}

	return q
}

func (q Query) QueryType() string {
	return q.queryType
}

func (q Query) ProjectID() string {
	return q.projectID
}
