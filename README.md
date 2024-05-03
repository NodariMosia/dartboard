# Dartboard

## About

TODO

## Setup

```bash
make setup

# or

go mod tidy
go mod download
```

## Usage

### Run

```bash
make run

# or

go run ./cmd
```

### Build

```bash
make build

# or

rimraf bin
go mod tidy
go build -o bin/ ./cmd
```

### Launch the built binary

```bash
make launch

# or

./bin/cmd
```
