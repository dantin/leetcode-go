
CURDIR := $(shell pwd)
GO := go

TEST_PKGS := $(shell find . -iname "*_test.go" -exec dirname {} \;)

.PHONY: build hamming_distance array_partition_i

hamming_distance:
	$(GO) build -o bin/hamming_distance cmd/hamming_distance/main.go

array_partition_i:
	$(GO) build -o bin/array_partition_i cmd/array_partition_i/main.go

test:
	# testing...
	$(GO) test -race -cover $(TEST_PKGS)


