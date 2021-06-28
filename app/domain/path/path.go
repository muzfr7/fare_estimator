package path

import "time"

// Path represents a position in a ride.
type Path struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Timestamp int32   `json:"timestamp"`
}

// IsDayTime checks whether its a day time.
func (p *Path) IsDayTime() bool {
	t := time.Unix(int64(p.Timestamp), 0).UTC()
	hour := t.Hour()
	if hour >= 5 && hour < 24 {
		return true
	}

	return false
}
