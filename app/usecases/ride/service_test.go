// +build unit

package ride_test

import (
	"reflect"
	"testing"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	rideDomain "github.com/muzfr7/fare_estimator/app/domain/ride"
	rideUsecase "github.com/muzfr7/fare_estimator/app/usecases/ride"
)

// TestCreate is a unit test for Create method.
func TestCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		row []string
	}

	tests := []struct {
		name         string
		args         args
		expectedRide rideDomain.Ride
		expectedErr  error
	}{
		{
			name: "Happy path",
			args: args{row: []string{"1", "37.964168", "23.726123", "1405595110"}},
			expectedRide: rideDomain.Ride{
				ID: 1,
				Paths: []pathDomain.Path{
					{
						Latitude:  37.964168,
						Longitude: 23.726123,
						Timestamp: 1405595110,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := rideUsecase.NewService()

			// input read only row channel
			rowChan := make(chan []string)
			go func() {
				defer close(rowChan)
				rowChan <- tc.args.row
			}()

			// output read only ride channel
			rideChan := make(chan rideDomain.Ride)
			go func() {
				defer close(rideChan)
				rideChan <- tc.expectedRide
			}()

			// output read only err channel
			errChan := make(chan error)
			go func() {
				defer close(errChan)
				errChan <- tc.expectedErr
			}()

			gotRideChan, gotErrChan := svc.Create(rowChan)
			if !reflect.DeepEqual(<-gotRideChan, <-rideChan) {
				t.Errorf("Create() got ride: %v, want ride: %v", <-gotRideChan, <-rideChan)
			}
			if !reflect.DeepEqual(<-gotErrChan, <-errChan) {
				t.Errorf("Create() got err: %v, want err: %v", <-gotErrChan, <-errChan)
			}
		})
	}
}
