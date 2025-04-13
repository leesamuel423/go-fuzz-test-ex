# go-fuzz-test-ex

Quick example on Go's built-in fuzzing capabilities.

## Requirements

- Go 1.18 or higher (for fuzzing support)

## Usage

### Running the Program

```bash
go run main.go # this will demonstrate string reversal fn w/ ex sentence
```

### Running Tests

```bash
# Run standard unit tests
go test

# Run fuzz tests until manually stopped (Ctrl+C)
go test -fuzz=Fuzz

# Run fuzz tests for a specific duration
go test -fuzz=Fuzz -fuzztime 30s
```

## How Fuzzing Works

The fuzzer automatically generates random input values to test the `Reverse` function. When it finds an input that causes unexpected behavior, it:
1. Fails the test
2. Reports the issue
3. Saves the failing input to the `testdata/fuzz` directory for future regression testing

The fuzz test in this project checks that:
- Reversing a string twice returns the original string
- The reversed string maintains UTF-8 validity
