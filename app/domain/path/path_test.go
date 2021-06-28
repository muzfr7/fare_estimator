// +build unit

package path_test

import (
	"testing"

	pathDomain "github.com/muzfr7/fare_estimator/app/domain/path"
	"github.com/stretchr/testify/assert"
)

// TestIsDayTime is a unit test for IsDayTime method.
func TestIsDayTime(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		path     pathDomain.Path
		expected bool
	}{
		{
			name:     "When day time is given",
			path:     pathDomain.Path{Timestamp: 1405594957},
			expected: true,
		},
		{
			name:     "When night time is given",
			path:     pathDomain.Path{Timestamp: 1624754392},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.path.IsDayTime()

			assert.Equal(t, tc.expected, got)
		})
	}
}
