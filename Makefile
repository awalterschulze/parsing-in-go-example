dependencies:
	go get -u github.com/goccmack/gocc

.PHONY: regenerate
regenerate:
	(cd example01_edge && make regenerate)

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build ./...

.PHONY: check
check:
	gofmt -l -s -w .
	git diff --exit-code
