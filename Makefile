build:
	go build ./...

test:
	go test -race ./...

fmt:
	gofmt -s -w .

vet:
	go vet ./...