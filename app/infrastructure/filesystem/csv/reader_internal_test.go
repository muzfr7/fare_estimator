// +build unit

package csv

import (
	"reflect"
	"testing"
)

// TestNewReader is a unit test for NewReader method.
func TestNewReader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want Reader
	}{
		{
			name: "Happy path",
			want: &readerImpl{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader(): %v, want: %v", got, tt.want)
			}
		})
	}
}
