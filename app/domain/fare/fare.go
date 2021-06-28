package fare

// Fare represents an estimated fare for a ride.
type Fare struct {
	RideID          uint64  `json:"id_ride"`
	EstimatedAmount float64 `json:"fare_estimate"`
}
