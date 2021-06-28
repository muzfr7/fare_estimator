package ride

import (
	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
)

// Ride represents a ride with its path.
type Ride struct {
	ID    uint64            `json:"id_ride"`
	Paths []pathDomain.Path `json:"paths"`
}
