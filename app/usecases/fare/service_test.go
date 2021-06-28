// +build unit

package fare_test

import (
	"reflect"
	"testing"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	rideDomain "github.com/muzfr7/fare_estimator/app/domain/ride"
	fareUsecase "github.com/muzfr7/fare_estimator/app/usecases/fare"
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
)

// TestEstimate is a unit test for Estimate method.
func TestEstimate(t *testing.T) {
	type fields struct {
		pathSVC pathUsecase.Service
	}

	type args struct {
		ride rideDomain.Ride
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		expectedFare []string
	}{
		{
			name: "Happy path",
			fields: fields{
				pathSVC: nil,
			},
			args: args{
				ride: rideDomain.Ride{
					ID: 1,
					Paths: []pathDomain.Path{
						{
							Latitude:  37.964168,
							Longitude: 23.726123,
							Timestamp: 1405595110,
						},
					},
				},
			},
			expectedFare: []string{"1", "3.47"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := fareUsecase.NewService(tc.fields.pathSVC)

			// input read only ride channel
			rideChan := make(chan rideDomain.Ride)
			go func() {
				defer close(rideChan)
				rideChan <- tc.args.ride
			}()

			// output read only fare channel
			fareChan := make(chan []string)
			go func() {
				defer close(fareChan)
				fareChan <- tc.expectedFare
			}()

			gotFareChan := svc.Estimate(rideChan)
			if !reflect.DeepEqual(<-gotFareChan, <-fareChan) {
				t.Errorf("Estimate(): %v, want: %v", <-gotFareChan, <-fareChan)
			}
		})
	}
}
