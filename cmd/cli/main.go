package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "", "-file paths.csv")
	flag.Parse()
	if *file == "" {
		fmt.Println("file parameter is required.")
		os.Exit(1)
	}
func mergeErrorChannels(channels ...<-chan error) <-chan error {
	errChan := make(chan error)

	for _, channel := range channels {
		go func(ch <-chan error) {
			for err := range ch {
				errChan <- err
			}
		}(channel)
	}

	return errChan
}
