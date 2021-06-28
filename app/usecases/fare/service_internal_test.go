// +build unit

package fare

import (
	"reflect"
	"testing"

	fareDomain "github.com/muzfr7/fare_estimator/app/domain/fare"
	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	rideDomain "github.com/muzfr7/fare_estimator/app/domain/ride"
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
)

// TestNewService is a unit test for NewService method.
func TestNewService(t *testing.T) {
	t.Parallel()

	type args struct {
		pathSVC pathUsecase.Service
	}

	testCases := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "Happy path",
			args: args{
				pathSVC: nil,
			},
			want: &serviceImpl{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NewService(tc.args.pathSVC); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NewService(): %v, want: %v", got, tc.want)
			}
		})
	}
}

// TestEstimateFor is a unit test for EstimateFor method.
func TestEstimateFor(t *testing.T) {
	t.Parallel()

	type fields struct {
		pathSVC pathUsecase.Service
	}

	type args struct {
		ride *rideDomain.Ride
	}

	testCases := []struct {
		name         string
		fields       fields
		args         args
		expectedFare *fareDomain.Fare
	}{
		{
			name: "Happy path",
			fields: fields{
				pathSVC: nil,
			},
			args: args{
				ride: &rideDomain.Ride{
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
			expectedFare: &fareDomain.Fare{
				RideID:          1,
				EstimatedAmount: 3.47,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &serviceImpl{
				pathSVC: tc.fields.pathSVC,
			}

			if got := svc.estimateFor(tc.args.ride); !reflect.DeepEqual(got, tc.expectedFare) {
				t.Errorf("estimateFor(): %v, want: %v", got, tc.expectedFare)
			}
		})
	}
}
