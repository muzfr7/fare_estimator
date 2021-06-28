GO ?= go

CLI_CMD = "cmd/cli/main.go"
CLI_BIN = "./bin/fareestimatorcli"

INPUT_CSV_FILE = "testdata/paths.csv"
RESULT_CSV_FILE = "./testdata/result.csv"

COVERAGE_PATH = "./..."
COVERAGE_FILE = "testdata/coverage.out"

MOCK_GENERATE_PATH = "./..."

# Remove generated binary, csv and coverage files
clean:
	@rm $(RESULT_CSV_FILE) $(CLI_BIN) $(COVERAGE_FILE)

# Generate mocks for interfaces
mocks:
	@$(GO) generate $(MOCK_GENERATE_PATH)

# Run tests
tests:
	@$(GO) test -count=1 -race -v -tags=unit --parallel 10 $(COVERAGE_PATH) -coverprofile=$(COVERAGE_FILE)

# View test coverage in browser
coverage:
	@$(GO) tool cover -html=$(COVERAGE_FILE)

# Build and run CLI
run:
	@$(GO) build -o $(CLI_BIN) $(CLI_CMD)
	@$(CLI_BIN) -file=$(INPUT_CSV_FILE)
