run:
	go run main.go

test:
	go test -v ./...

lint:
	gofumpt -w -s ./..
	golangci-lint run --fix

