deps:
	go get -t -v ./...
	go get -u golang.org/x/tools/cmd/vet
	go get -u golang.org/x/tools/cmd/cover

test:
	go vet ./...
	go test -cover -short ./...

.PHONY: deps, test