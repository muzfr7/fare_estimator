// +build unit

package ride

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		{
			name: "Happy path",
			want: &serviceImpl{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
