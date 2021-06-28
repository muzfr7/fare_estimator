// +build unit

package fare

import (
	"reflect"
	"testing"

	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
)

// TestNewService is a unit test for NewService method.
func TestNewService(t *testing.T) {
	t.Parallel()

	type args struct {
		pathSVC pathUsecase.Service
	}

	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.pathSVC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService(): %v, want: %v", got, tt.want)
			}
		})
	}
}
