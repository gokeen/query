package query

// Count resource
// https://keen.io/docs/api/reference/#count-resource
func Count(coll string, opts ...queryOpts) *Query {
	return new("count", coll, "", opts...)
}

// CountUnique resource
// https://keen.io/docs/api/reference/#count-unique-resource
func CountUnique(coll, targ string, opts ...queryOpts) *Query {
	return new("count_unique", coll, targ, opts...)
}

// Minimum resource
// https://keen.io/docs/api/reference/#minimum-resource
func Minimum(coll, targ string, opts ...queryOpts) *Query {
	return new("minimum", coll, targ, opts...)
}

// Maximum resource
// https://keen.io/docs/api/reference/#maximum-resource
func Maximum(coll, targ string, opts ...queryOpts) *Query {
	return new("maximum", coll, targ, opts...)
}

// Average resource
// https://keen.io/docs/api/reference/#average-resource
func Average(coll, targ string, opts ...queryOpts) *Query {
	return new("average", coll, targ, opts...)
}

// Median resource
// https://keen.io/docs/api/reference/#median-resource
func Median(coll, targ string, opts ...queryOpts) *Query {
	return new("median", coll, targ, opts...)
}

// Percentile resource
// https://keen.io/docs/api/reference/#percentile-resource
// This will require all percentiles (p) as 2 decimal floats, eg. 95.00 not 95
func Percentile(coll, targ string, p float64, opts ...queryOpts) *Query {
	q := new("percentile", coll, targ, opts...)
	q.Percentile = p
	return q
}

// Sum resource
// https://keen.io/docs/api/reference/#sum-resource
func Sum(coll, targ string, opts ...queryOpts) *Query {
	return new("sum", coll, targ, opts...)
}

// SelectUnique resource
// https://keen.io/docs/api/reference/#select-unique-resource
func SelectUnique(coll, targ string, opts ...queryOpts) *Query {
	return new("select_unique", coll, targ, opts...)
}

// Extraction resource
// https://keen.io/docs/api/reference/#extraction-resource
func Extraction(coll string, opts ...queryOpts) *Query {
	return new("extraction", coll, "", opts...)
}
