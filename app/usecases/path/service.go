//go:generate mockery --name Service

package path

import (
	"math"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
)

const (
	earthRadius = float64(6371)
)

// Service provides required methods to manipulate ride path.
type Service interface {
	CalculateDistance(startPath, endPath pathDomain.Path) float64
}

// serviceImpl implements Service interface.
type serviceImpl struct{}

// NewService returns a new instance of serviceImpl.
func NewService() Service {
	return &serviceImpl{}
}

// CalculateDistance uses the ‘haversine’ formula to calculate the great-circle distance between two points
// that is, the shortest distance over the earth’s surface
// giving an ‘as-the-crow-flies’ distance between the points (ignoring any hills they fly over, of course!).
func (s *serviceImpl) CalculateDistance(startPath, endPath pathDomain.Path) float64 {
	// distance between latitudes and longitudes
	var deltaLat = (endPath.Latitude - startPath.Latitude) * (math.Pi / 180)
	var deltaLong = (endPath.Longitude - startPath.Longitude) * (math.Pi / 180)

	// convert to radians
	lat1 := startPath.Latitude * (math.Pi / 180)
	lat2 := endPath.Latitude * (math.Pi / 180)

	// apply formulae
	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLong/2)*math.Sin(deltaLong/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
