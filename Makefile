echo:
	echo "Hello World!"

setup:
	go version
	go mod tidy
	go mod download

run:
	go run ./cmd

clean:
	rimraf bin
	go clean
	go mod tidy
	go mod download

build:
	rimraf bin
	go mod tidy
	go build -o bin/ ./cmd

launch:
	./bin/cmd
