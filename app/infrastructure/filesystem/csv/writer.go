//go:generate mockery --name Writer

package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

const (
	outputFile = "testdata/result.csv"
)

// Writer provides required methods to write csv.
type Writer interface {
	Write(in <-chan []string) (<-chan int, <-chan error, error)
}

// writerImpl implements Writer interface.
type writerImpl struct{}

// NewWriter returns a new instance of writerImpl.
func NewWriter() Writer {
	return &writerImpl{}
}

// Write will write out estimated fares for each ride in a new csv file.
func (w *writerImpl) Write(fareChan <-chan []string) (<-chan int, <-chan error, error) {
	// create absolute path from given path
	fullpath, err := filepath.Abs(outputFile)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid file path: %v; error: %v ", outputFile, err)
	}

	// get parent directory path
	dirpath := filepath.Dir(fullpath)

	// set parent dir permissions to 0777
	err = os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create path: %v; error: %v ", outputFile, err)
	}

	// create file
	file, err := os.Create(fullpath)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create file; error: %v", err)
	}

	csvWriter := csv.NewWriter(file)

	// create channels for status and error
	doneChan := make(chan int)
	errChan := make(chan error)

	go func() {
		// write each ride fare row into the file
		for row := range fareChan {
			err := csvWriter.Write(row)
			if err != nil {
				errChan <- fmt.Errorf("cannot write row in file; error: %v", err)
			}
		}

		csvWriter.Flush()

		// close file
		file.Close()

		doneChan <- 0

		// close channels
		close(doneChan)
		close(errChan)
	}()

	return doneChan, errChan, nil
}
