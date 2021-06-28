// +build unit

package csv

import (
	"reflect"
	"testing"
)

// TestNewReader is a unit test for NewReader method.
func TestNewReader(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		want Reader
	}{
		{
			name: "Happy path",
			want: &readerImpl{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NewReader(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NewReader(): %v, want: %v", got, tc.want)
			}
		})
	}
}
