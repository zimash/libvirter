libvirter:
	go build -o ./bin/libvirter cmd/cli/libvirter.go

clean:
	rm -rf bin

build: clean libvirter

test:
	go test ./...

lint:
	golangci-lint run -v

