# Fare Estimator
A fare estimator app which receives cab positions/paths for each ride in a csv file and returns a csv file with estimated fares for each ride. It uses `Haversine distance formula` to calculate distance between two (lat, lng) pairs.

## Constraints
```
Input paths.csv consist of:
    id_ride, lat, lng, timestamp

Output result.csv will contain:
    id_ride, fare_estimate

Discard invalid entries:
    Filtering should be performed as follows: consecutive tuples p1, p2 should be used to calculate the segment’s speed U. If U > 100km/h, p2 should be removed from the set.

Fare estimation:

Amount when cab is moving at > 10km/h during day time (05:00, 00:00)
    0.74 per km

Amount when cab is moving at > 10km/h during night time (00:00, 05:00)
    1.30 per km

Amount when cab is idle <= 10km/h at anytime
    11.90 per hour of idle time

Standard flag amount is charged to the rides fare
    1.30, the minimum ride fare should be at least 3.47
```

## Setup

Clone this repository somewhere on your computer
```bash
git clone git@github.com:muzfr7/fare_estimator.git
```

> make sure you have `paths.csv` file in `testdata` directory, because it will be ingested directly for fare estimation.

## Usage

Before proceeding further, make sure to change directory into cloned `fare_estimator`

### Build and Run
```bash
make run
```
> this will build and run `./bin/fareestimatorcli` ingesting `./testdata/paths.csv` file, to produce `./testdata/result.csv` file containing estimated fares for each ride.

### Run tests
```bash
make tests
```
> this will also generate `./testdata/coverage.out` file, issue a `make coverage` command to view it in browser.

### Cleanup
Issue following command to remove all generated files: `./bin/fareestimatorcli`, `./testdata/result.csv` and `./testdata/coverage.out`
```bash
make clean
```
