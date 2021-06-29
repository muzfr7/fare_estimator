// +build unit

package main

import (
	"reflect"
	"testing"
)

// TestMergeErrorChannels is a unit test for mergeErrorChannels method.
func TestMergeErrorChannels(t *testing.T) {
	t.Parallel()

	type args struct {
		channels error
	}

	testCases := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "Happy path",
			args: args{
				channels: nil,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// input read only channels
			channels := make(chan error)
			go func() {
				defer close(channels)
				channels <- tc.args.channels
			}()

			// output read only err channel
			errChan := make(chan error)
			go func() {
				defer close(errChan)
				errChan <- tc.expectedErr
			}()

			gotErrChan := mergeErrorChannels(channels)
			if !reflect.DeepEqual(<-gotErrChan, <-errChan) {
				t.Errorf("mergeErrorChannels(): %v, want: %v", <-gotErrChan, <-errChan)
			}
		})
	}
}
