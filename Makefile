libvirter:
	go build -o ./bin/libvirter cmd/cli/libvirter.go

clean:
	rm -rf bin
