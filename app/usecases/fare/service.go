//go:generate mockery --name Service

package fare

import (
	"fmt"
	"math"

	fareDomain "github.com/muzfr7/fare_estimator/app/domain/fare"
	rideDomain "github.com/muzfr7/fare_estimator/app/domain/ride"
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
)

const (
	buffSize = 50
)

const (
	minSpeed             = float64(10)
	maxSpeed             = float64(100)
	flagRate             = float64(1.3)
	idleHourRate         = float64(11.9)
	minTotal             = float64(3.47)
	movingRateDayShift   = float64(0.74)
	movingRateNightShift = float64(1.3)
)

// Service provides required methods to manipulate ride fare.
type Service interface {
	Estimate(rideChan <-chan rideDomain.Ride) <-chan []string
}

// serviceImpl implements Service interface.
type serviceImpl struct {
	pathSVC pathUsecase.Service
}

// NewService returns a new instance of serviceImpl.
func NewService(pathSVC pathUsecase.Service) Service {
	return &serviceImpl{
		pathSVC: pathSVC,
	}
}

// Estimate will estimate fare for rides.
func (s *serviceImpl) Estimate(rideChan <-chan rideDomain.Ride) <-chan []string {
	// create channel for fare
	fareChan := make(chan []string, buffSize)

	go func() {
		for ride := range rideChan {
			estimatedFare := s.estimateFor(&ride)

			fareChan <- []string{
				fmt.Sprintf("%v", estimatedFare.RideID),
				fmt.Sprintf("%.2f", estimatedFare.EstimatedAmount),
			}
		}

		// close channel
		close(fareChan)
	}()

	return fareChan
}

func (s *serviceImpl) estimateFor(ride *rideDomain.Ride) *fareDomain.Fare {
	// standard `flag` amount
	total := flagRate

	// O(n²), must be reduced to O(n) at-least
	for i := 0; i < len(ride.Paths)-1; i++ {
		for j := i + 1; j < len(ride.Paths); j++ {

			startPath := ride.Paths[i]
			endPath := ride.Paths[i+1]

			// elapsed time in seconds between start and end paths
			deltaTimeInSeconds := float64(endPath.Timestamp - startPath.Timestamp)

			// distance between start and end paths using Haversine formula
			deltaDistanceInKM := s.pathSVC.CalculateDistance(startPath, endPath)

			speedKMH := (deltaDistanceInKM / deltaTimeInSeconds) * 3600

			// discard invalid path entry
			if speedKMH > maxSpeed {
				i++
				continue
			}

			// calculate idle rate
			if speedKMH <= minSpeed {
				total += (deltaTimeInSeconds / 3600) * idleHourRate
				break
			}

			// calculate distance rate by hour
			if startPath.IsDayTime() {
				// day time
				total += deltaDistanceInKM * movingRateDayShift
			} else {
				// night time
				total += deltaDistanceInKM * movingRateNightShift
			}

			break
		}
	}

	total = math.Max(total, minTotal)

	return &fareDomain.Fare{
		RideID:          ride.ID,
		EstimatedAmount: total,
	}
}
