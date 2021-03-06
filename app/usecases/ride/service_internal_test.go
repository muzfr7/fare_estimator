// +build unit

package ride

import (
	"fmt"
	"reflect"
	"testing"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	"github.com/stretchr/testify/assert"
)

// TestNewService is a unit test for NewService method.
func TestNewService(t *testing.T) {
	t.Parallel()

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
				t.Errorf("NewService(): %v, want: %v", got, tt.want)
			}
		})
	}
}

// TestParseRow is a unit test for ParseRow method.
func TestParseRow(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		row           []string
		expectedID    uint64
		expectedPath  *pathDomain.Path
		expectedError error
	}{
		{
			name:          "Happy path",
			row:           []string{"1", "37.964168", "23.726123", "1405595110"},
			expectedID:    1,
			expectedPath:  &pathDomain.Path{Latitude: 37.964168, Longitude: 23.726123, Timestamp: 1405595110},
			expectedError: nil,
		},
		{
			name:          "When row doesn't contain 4 elements",
			row:           []string{"2", "35.355555"},
			expectedError: fmt.Errorf("row doesn't contain 4 elements: %v", []string{"2", "35.355555"}),
		},
		{
			name:          "When failed to convert ride id string into uint",
			row:           []string{"abc", "37.964168", "23.726123", "1405595110"},
			expectedError: fmt.Errorf("failed to convert ride id string into uint"),
		},
		{
			name:          "When failed to convert latitude string into float",
			row:           []string{"3", "abc", "23.726123", "1405595110"},
			expectedError: fmt.Errorf("failed to convert latitude string into float"),
		},
		{
			name:          "When failed to convert longitude string into float",
			row:           []string{"3", "37.964168", "abc", "1405595110"},
			expectedError: fmt.Errorf("failed to convert longitude string into float"),
		},
		{
			name:          "When failed to convert timestamp string into int",
			row:           []string{"3", "37.964168", "23.726123", "abc"},
			expectedError: fmt.Errorf("failed to convert timestamp string into int"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, path, err := parseRow(tc.row)

			if err == nil {
				assert.Equal(t, tc.expectedID, id)
				assert.Equal(t, tc.expectedPath, path)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, tc.expectedError, err.Error())
			}
		})
	}
}
