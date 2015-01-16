package query

// Filter represents the filter sturcture for queries
type Filter struct {
	PropertyName  string      `json:"property_name"`
	Operator      string      `json:"operator"`
	PropertyValue interface{} `json:"property_value"`
}

// GeoLoc is a sub type for PropertyValue field in Filter for Geo Filtering
// Coordinates should always be in [<longitude>, <latitude>] order
type GeoLoc struct {
	Coordinates      [2]float64 `json:"coordinates"`
	MaxDistanceMiles int        `json:"max_distance_miles,omitempty"`
}

// NewGeo returns a new GeoLoc
func NewGeo(lon, lat float64, opts ...func(g *GeoLoc)) *GeoLoc {
	g := &GeoLoc{
		Coordinates: [2]float64{lon, lat},
	}
	for _, v := range opts {
		v(g)
	}
	return g
}
