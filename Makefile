
CURDIR := $(shell pwd)
GO := go

TEST_PKGS := $(shell find . -iname "*_test.go" -exec dirname {} \;)

.PHONY: build hamming_distance

hamming_distance:
	$(GO) build -o bin/hamming_distance cmd/hamming_distance/main.go

test:
	# testing...
	$(GO) test -race -cover $(TEST_PKGS)


