# Dartboard

## About

This project is a golang TUI (Terminal User Interface) dartboard game.

Project uses [bubbletea](https://github.com/charmbracelet/bubbletea) to build the
terminal user interface.

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
