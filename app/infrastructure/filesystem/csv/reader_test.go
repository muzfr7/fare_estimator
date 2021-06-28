// +build unit

package csv_test

import (
	"reflect"
	"testing"

	appCSV "github.com/muzfr7/fare_estimator/app/infrastructure/filesystem/csv"
)

// TestRead is a unit test for Read method.
func TestRead(t *testing.T) {
	t.Parallel()

	type args struct {
		filePath string
	}

	testCases := []struct {
		name         string
		args         args
		expectedRow  []string
		expectedErr1 error
		expectedErr2 bool
	}{
		{
			name: "Happy path",
			args: args{
				filePath: "../../../../testdata/paths_for_test.csv",
			},
			expectedRow:  []string{"1", "37.966660", "23.728308", "1405594957"},
			expectedErr1: nil,
			expectedErr2: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := appCSV.NewReader()

			// output read only row channel
			rowChan := make(chan []string)
			go func() {
				defer close(rowChan)
				rowChan <- tc.expectedRow
			}()

			// output read only err channel
			errChan := make(chan error)
			go func() {
				defer close(errChan)
				errChan <- tc.expectedErr1
			}()

			gotRowChan, gotErrChan, err := reader.Read(tc.args.filePath)
			if (err != nil) != tc.expectedErr2 {
				t.Errorf("Read() err: %v, want err: %v", err, tc.expectedErr2)
				return
			}
			if !reflect.DeepEqual(<-gotRowChan, <-rowChan) {
				t.Errorf("Read() got row: %v, want: %v", <-gotRowChan, <-rowChan)
			}
			if !reflect.DeepEqual(<-gotErrChan, <-errChan) {
				t.Errorf("Read() got err: %v, want err: %v", <-gotErrChan, <-errChan)
			}
		})
	}
}
