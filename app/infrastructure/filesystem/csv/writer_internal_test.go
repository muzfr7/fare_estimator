// +build unit

package csv

import (
	"reflect"
	"testing"
)

// TestNewWriter is a unit test for NewWriter method.
func TestNewWriter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want Writer
	}{
		{
			name: "Happy path",
			want: &writerImpl{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWriter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter(): %v, want: %v", got, tt.want)
			}
		})
	}
}
