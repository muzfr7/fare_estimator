package main

import (
	"flag"
	"fmt"
	"os"

	appCSV "github.com/muzfr7/fare_estimator/app/infrastructure/filesystem/csv"
	fareUsecase "github.com/muzfr7/fare_estimator/app/usecases/fare"
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
	rideUsecase "github.com/muzfr7/fare_estimator/app/usecases/ride"
	logger "github.com/sirupsen/logrus"
)

func main() {
	file := flag.String("file", "", "-file paths.csv")
	flag.Parse()
	if *file == "" {
		fmt.Println("file parameter is required.")
		os.Exit(1)
	}

	csvReader := appCSV.NewReader()
	csvWriter := appCSV.NewWriter()

	rideSVC := rideUsecase.NewService()
	pathSVC := pathUsecase.NewService()
	fareSVC := fareUsecase.NewService(pathSVC)

	if err := handler(*file, csvReader, csvWriter, rideSVC, fareSVC); err != nil {
		panic(err)
	}

	logger.Println("Fare estimation completed, please check `testdata/result.csv` file..")
}

func handler(file string, csvReader appCSV.Reader, csvWriter appCSV.Writer, rideSVC rideUsecase.Service, fareSVC fareUsecase.Service) error {
	// read each row of csv file in rowChan
	rowChan, rowChanErr, err := csvReader.Read(file)
	if err != nil {
		return fmt.Errorf("failed to read file; error: %v", err)
	}

	// create ride.Ride from rowChan and return it in rideChan
	rideChan, rideChanErr := rideSVC.Create(rowChan)

	// estimate fare from rideChan for each ride and return it in fareChan
	fareChan := fareSVC.Estimate(rideChan)

	// create a new csv file with estimated ride fares from fareChan
	done, writeChanErr, err := csvWriter.Write(fareChan)
	if err != nil {
		return fmt.Errorf("failed to initialize write channel; error: %v", err)
	}
	select {
	case <-done:
	case err := <-mergeErrorChannels(rowChanErr, rideChanErr, writeChanErr):
		if err != nil {
			return fmt.Errorf("failed to perform estimation; received an error from channels; error: %v", err)
		}
	}

	return nil
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
