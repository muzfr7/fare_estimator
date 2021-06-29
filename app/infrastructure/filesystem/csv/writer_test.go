// +build unit

package csv_test

import (
	"reflect"
	"testing"

	appCSV "github.com/muzfr7/fare_estimator/app/infrastructure/filesystem/csv"
)

const (
	testOutputFile = "../../../../testdata/result_for_test.csv"
)

// TestWrite is a unit test for Write method.
func TestWrite(t *testing.T) {
	type args struct {
		fare []string
	}

	testCases := []struct {
		name         string
		args         args
		expectedDone int
		expectedErr1 error
		expectedErr2 bool
	}{
		{
			name: "Happy path",
			args: args{
				fare: []string{"1", "3.47"},
			},
			expectedDone: 0,
			expectedErr1: nil,
			expectedErr2: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			writer := appCSV.NewWriter(testOutputFile)

			// input read only fare channel
			fareChan := make(chan []string)
			go func() {
				defer close(fareChan)
				fareChan <- tc.args.fare
			}()

			// output read only done channel
			doneChan := make(chan int)
			go func() {
				defer close(doneChan)
				doneChan <- tc.expectedDone
			}()

			// output read only err channel
			errChan := make(chan error)
			go func() {
				defer close(errChan)
				errChan <- tc.expectedErr1
			}()

			gotDoneChan, gotErrChan, err := writer.Write(fareChan)
			if (err != nil) != tc.expectedErr2 {
				t.Errorf("Write() error: %v, want err: %v", err, tc.expectedErr2)
				return
			}
			if !reflect.DeepEqual(<-gotDoneChan, <-doneChan) {
				t.Errorf("Write() got done: %v, want done: %v", <-gotDoneChan, <-doneChan)
			}
			if !reflect.DeepEqual(<-gotErrChan, <-errChan) {
				t.Errorf("Write() got err: %v, want err: %v", <-gotErrChan, <-errChan)
			}
		})
	}
}
