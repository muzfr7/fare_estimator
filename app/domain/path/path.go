package path

// Path represents a position in a ride.
type Path struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Timestamp int32   `json:"timestamp"`
}
