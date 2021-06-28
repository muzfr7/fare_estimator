//go:generate mockery --name Reader

package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	buffSize = 50
)

// Reader provides required methods to parse csv.
type Reader interface {
	Read(filePath string) (<-chan []string, <-chan error, error)
}

// readerImpl implements Reader interface.
type readerImpl struct{}

// NewReader returns a new instance of readerImpl.
func NewReader() Reader {
	return &readerImpl{}
}

// Read will read the file line by line and returns a []string channel.
func (r *readerImpl) Read(filePath string) (<-chan []string, <-chan error, error) {
	// create absolute path from given path
	fullpath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid file path: %v; error: %v ", filePath, err)
	}

	// open the file for reading
	file, err := os.Open(fullpath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v; error: %v ", fullpath, err)
	}

	csvReader := csv.NewReader(bufio.NewReader(file))

	// create channels for row and error
	rowChan := make(chan []string, buffSize)
	errChan := make(chan error)

	// spawn a goroutine to concurrently read data from file
	go func() {
		// read each row from the file
		for {
			row, err := csvReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				errChan <- fmt.Errorf("failed to read from file: %v; error: %v ", fullpath, err)
			}

			rowChan <- row
		}

		// close file
		file.Close()

		// close channels
		close(rowChan)
		close(errChan)
	}()

	return rowChan, errChan, nil
}
