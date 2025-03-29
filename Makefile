GO := go
DIST := bin

BINARY = $(DIST)/monkey

SRCS := $(shell find . -type f -name "*.go")

$(warning SRCS = $(SRCS))

.PHONY: run


$(BINARY): $(SRCS)
	go build -o $@ .

build: $(BINARY)

run: build
	@./$(BINARY)

test:
	go test ./... -v

