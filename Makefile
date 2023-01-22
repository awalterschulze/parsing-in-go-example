dependencies:
	go get -u github.com/goccmack/gocc
	go install github.com/goccmack/gocc

.PHONY: regenerate
regenerate:
	(cd example01_edge && make regenerate)
	(cd example02_sdt && make regenerate)
	(cd example03_edges && make regenerate)
	make fmt

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build ./...

.PHONY: fmt
fmt:
	gofmt -l -s -w .

.PHONY: check
check:
	make fmt
	git diff --exit-code
