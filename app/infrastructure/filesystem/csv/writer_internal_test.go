// +build unit

package csv

import (
	"reflect"
	"testing"
)

// TestNewWriter is a unit test for NewWriter method.
func TestNewWriter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		want Writer
	}{
		{
			name: "Happy path",
			want: &writerImpl{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NewWriter(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NewWriter(): %v, want: %v", got, tc.want)
			}
		})
	}
}
