// +build unit

package path_test

import (
	"testing"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
)

// TestCalculateDistance is a unit test for CalculateDistance method.
func TestCalculateDistance(t *testing.T) {
	type args struct {
		startPath pathDomain.Path
		endPath   pathDomain.Path
	}

	testCases := []struct {
		name             string
		args             args
		expectedDistance float64
	}{
		{
			name: "Happy path",
			args: args{
				startPath: pathDomain.Path{
					Latitude:  37.966660,
					Longitude: 23.728308,
				},
				endPath: pathDomain.Path{
					Latitude:  37.935490,
					Longitude: 23.625655,
				},
			},
			expectedDistance: 9.645003763681704,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svc := pathUsecase.NewService()

			if got := svc.CalculateDistance(tc.args.startPath, tc.args.endPath); got != tc.expectedDistance {
				t.Errorf("CalculateDistance(): %v, want: %v", got, tc.expectedDistance)
			}
		})
	}
}
