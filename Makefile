.PHONY: all
all: build

.PHONY: deps
deps:
	go get .

.PHONY: tables
tables:
	go generate .

.PHONY: build
build: deps
	go build -o color cmd/color/main.go

.PHONY: install
install: deps
	go install cmd/color/main.go

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test -v ./...

.PHONY: tag
tag:
	sh scripts/bumpup.sh

