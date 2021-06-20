.PHONY: all
all: build

.PHONY: deps
deps:
	go get -v -d

.PHONY: build
build: deps
	go build -o color cmd/color/color.go

.PHONY: install
install: deps
	go install cmd/color/color.go

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test -v .

.PHONY: tag
tag:
	sh scripts/bumpup.sh

.PHONY: tables
tables:
	sh scripts/create-tables.sh
