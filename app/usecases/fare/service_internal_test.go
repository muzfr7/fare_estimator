// +build unit

package fare

import (
	"reflect"
	"testing"

	fareDomain "github.com/muzfr7/fare_estimator/app/domain/fare"
	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	rideDomain "github.com/muzfr7/fare_estimator/app/domain/ride"
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
	pathUsecaseMocks "github.com/muzfr7/fare_estimator/app/usecases/path/mocks"
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

	pathSVCMock := new(pathUsecaseMocks.Service)

	testCases := []struct {
		name             string
		fields           fields
		args             args
		expectedDistance float64
		expectedFare     *fareDomain.Fare
	}{
		{
			name: "Happy path",
			fields: fields{
				pathSVC: pathSVCMock,
			},
			args: args{
				ride: &rideDomain.Ride{
					ID: 1,
					Paths: []pathDomain.Path{
						{
							Latitude:  37.966660,
							Longitude: 23.728308,
							Timestamp: 1405594957,
						},
						{
							Latitude:  37.935490,
							Longitude: 23.625655,
							Timestamp: 1405596220,
						},
					},
				},
			},
			expectedDistance: 9.12,
			expectedFare: &fareDomain.Fare{
				RideID:          1,
				EstimatedAmount: 8.0488,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &serviceImpl{
				pathSVC: tc.fields.pathSVC,
			}

			startPath := tc.args.ride.Paths[0]
			endPath := tc.args.ride.Paths[1]

			// set expectations
			pathSVCMock.On("CalculateDistance", startPath, endPath).Return(tc.expectedDistance)

			if got := svc.estimateFor(tc.args.ride); !reflect.DeepEqual(got, tc.expectedFare) {
				t.Errorf("estimateFor(): %v, want: %v", got, tc.expectedFare)
			}
		})
	}
}
