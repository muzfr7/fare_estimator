//go:generate mockery --name Service

package ride

import (
	"fmt"
	"strconv"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	rideDomain "github.com/muzfr7/fare_estimator/app/domain/ride"
)

const (
	buffSize = 50
)

// Service provides required methods to manipulate ride.
type Service interface {
	Create(dataChan <-chan []string) (<-chan rideDomain.Ride, <-chan error)
}

// serviceImpl implements Service interface.
type serviceImpl struct{}

// NewService returns a new instance of serviceImpl.
func NewService() Service {
	return &serviceImpl{}
}

// Create will populate and return ride.Ride channel.
func (s *serviceImpl) Create(rowChan <-chan []string) (<-chan rideDomain.Ride, <-chan error) {
	// create channels for ride and error
	rideChan := make(chan rideDomain.Ride, buffSize)
	errChan := make(chan error)

	go func() {
		var rideID uint64
		var ridePaths []pathDomain.Path

		for row := range rowChan {
			// parse current row
			id, path, err := parseRow(row)
			if err != nil {
				errChan <- fmt.Errorf("failed to parse row; error: %v", err)
			}

			if len(ridePaths) != 0 && rideID != id {
				rideChan <- rideDomain.Ride{
					ID:    rideID,
					Paths: ridePaths,
				}

				// reset ridePaths
				ridePaths = []pathDomain.Path{}
			}

			rideID = id
			ridePaths = append(ridePaths, *path)
		}

		rideChan <- rideDomain.Ride{
			ID:    rideID,
			Paths: ridePaths,
		} // handles the final row

		// close channels
		close(rideChan)
		close(errChan)
	}()

	return rideChan, errChan
}

func parseRow(row []string) (uint64, *pathDomain.Path, error) {
	if len(row) < 4 {
		return 0, nil, fmt.Errorf("row doesn't contain 4 elements: %v", row)
	}

	id, err := strconv.ParseUint(row[0], 10, 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert ride id string into uint")
	}

	latitude, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert latitude string into float")
	}

	longitude, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert longitude string into float")
	}

	timestamp, err := strconv.ParseInt(row[3], 10, 32)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert timestamp string into int")
	}

	return id, &pathDomain.Path{
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: int32(timestamp),
	}, nil
}
