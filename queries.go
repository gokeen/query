package query

// Count resource
// https://keen.io/docs/api/reference/#count-resource
func Count(pid, coll string, opts ...queryOpts) *Query {
	return new(pid, "count", coll, "", opts...)
}

// CountUnique resource
// https://keen.io/docs/api/reference/#count-unique-resource
func CountUnique(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "count_unique", coll, targ, opts...)
}

// Minimum resource
// https://keen.io/docs/api/reference/#minimum-resource
func Minimum(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "minimum", coll, targ, opts...)
}

// Maximum resource
// https://keen.io/docs/api/reference/#maximum-resource
func Maximum(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "maximum", coll, targ, opts...)
}

// Average resource
// https://keen.io/docs/api/reference/#average-resource
func Average(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "average", coll, targ, opts...)
}

// Median resource
// https://keen.io/docs/api/reference/#median-resource
func Median(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "median", coll, targ, opts...)
}

// Percentile resource
// https://keen.io/docs/api/reference/#percentile-resource
// This will require all percentiles (p) as 2 decimal floats, eg. 95.00 not 95
func Percentile(pid, coll, targ string, p float64, opts ...queryOpts) *Query {
	q := new(pid, "percentile", coll, targ, opts...)
	q.Percentile = p
	return q
}

// Sum resource
// https://keen.io/docs/api/reference/#sum-resource
func Sum(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "sum", coll, targ, opts...)
}

// SelectUnique resource
// https://keen.io/docs/api/reference/#select-unique-resource
func SelectUnique(pid, coll, targ string, opts ...queryOpts) *Query {
	return new(pid, "select_unique", coll, targ, opts...)
}

// Extraction resource
// https://keen.io/docs/api/reference/#extraction-resource
func Extraction(pid, coll string, opts ...queryOpts) *Query {
	return new(pid, "extraction", coll, "", opts...)
}
