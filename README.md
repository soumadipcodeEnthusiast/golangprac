golangprac

this is practice repo for golanguage
Table of contents

    About
    Features
    Prerequisites
    Getting started
    Usage
    Project layout
    Testing
    Formatting & linting
    Contributing
    License
    Contact

About

golangprac is a small repository for practicing Go (Golang) — experimenting with language features, small utilities, examples, and learning idiomatic Go patterns.
Features

    Small, focused Go examples
    Playground code for learning and experimentation
    Tests and small CLI examples
    Opinionated project layout to practice real-world structure

Prerequisites

    Go 1.20+ installed (GOMAXPROCS and modules enabled)
    git (to clone the repo)

Getting started

    Clone the repo git clone https://github.com/soumadipcodeEnthusiast/golangprac.git cd golangprac

    Ensure modules are set up go mod tidy

    Build or run an example

        Run a simple package (replace ./cmd/example with the actual path you want to run): go run ./cmd/example

        Build a binary: go build -o bin/example ./cmd/example

Usage examples

    Run a package: go run ./path/to/package

    Run tests: go test ./...

    Run a specific test: go test ./pkg/somepkg -run TestSomething

    Run with race detector (helpful while learning concurrency): go test -race ./...

Project layout (suggested)

    cmd/ - applications (one folder per app)
    internal/ - internal packages not meant for public use
    pkg/ - reusable libraries intended to be importable
    examples/ - runnable examples & small programs
    tests/ - helper test data or integration tests
    go.mod, go.sum - module files
    README.md - this file

Adjust layout as you add experiments.
Testing

    Run unit tests: go test ./...

    Run with coverage: go test ./... -coverprofile=coverage.out go tool cover -html=coverage.out

Formatting & linting

    Format code: go fmt ./...

    Vet: go vet ./...

    Suggested linters (install separately):
        golangci-lint: https://github.com/golangci/golangci-lint

Contributing

Contributions, experiments, and issues are welcome. Keep changes small and focused (this repo is for learning). When opening a PR:

    Explain what you tried to learn or demonstrate.
    Add tests when appropriate.
    Keep code idiomatic and documented.

License

Choose a license for this repo (e.g., MIT). If you want, add a LICENSE file with your preferred license.
Contact

Owner: @soumadipcodeEnthusiast
