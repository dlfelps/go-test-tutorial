# Go Test Tutorial CLI Tool

A command-line educational tool for teaching beginners how to use Go's testing framework through interactive tutorials and examples.

## Features

- **Learn Mode**: Read about Go testing concepts step-by-step
- **Generate Mode**: Create test templates for your Go code
- **Run Mode**: Execute and analyze Go tests with helpful explanations
- **Interactive Tutorial**: Walk through all Go testing concepts interactively

## Installation

```bash
# Build the executable
go build -o gotest-learn

# Make it available in your path (optional)
cp gotest-learn /usr/local/bin/
```

## Usage

### Learning about Go testing

```bash
# List all available topics
./gotest-learn learn

# Read about a specific topic (e.g., topic 1: Introduction)
./gotest-learn learn 1

# Start interactive tutorial mode
./gotest-learn learn --interactive
```

### Generating test templates

```bash
# List available templates
./gotest-learn generate

# Generate a basic test file
./gotest-learn generate --type basic --file mycode

# Generate a table-driven test file
./gotest-learn generate --type table --file mycode

# Generate a benchmark test file  
./gotest-learn generate --type benchmark --file mycode
```

### Running tests

```bash
# Run all tests in current directory
./gotest-learn run

# Run tests in a specific file
./gotest-learn run --file mycode_test.go

# Run tests in a specific directory
./gotest-learn run --dir ./mypackage

# Run tests in verbose mode
./gotest-learn run --verbose
```

## Examples

The `examples` directory contains practical examples of Go testing techniques:

- **Calculator Package**: Demonstrates basic tests, table-driven tests, and benchmarks
  - Basic tests in `calculator_test.go`
  - Table-driven tests in `calculator_test.go`
  - Benchmarks in `calculator_bench_test.go`
  - Test helpers in `helpers_test.go`

## Topics Covered

1. Introduction to Go Testing
2. Writing Basic Tests
3. Test Assertions and Comparisons
4. Table-Driven Tests
5. Benchmarks
6. Subtests and Test Organization
7. TestMain and Setup/Teardown
8. Test Helpers and Utilities

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues for improvements or bug fixes.

## License

MIT