
CURDIR := $(shell pwd)
GO := go

TEST_PKGS := $(shell find . -iname "*_test.go" -exec dirname {} \;)

.PHONY: build hamming_distance array_partition_i number_complement \
	keyboard_row


build: hamming_distance array_partition_i number_complement keyboard_row

hamming_distance:
	$(GO) build -o bin/hamming_distance cmd/hamming_distance/main.go

array_partition_i:
	$(GO) build -o bin/array_partition_i cmd/array_partition_i/main.go

number_complement:
	$(GO) build -o bin/number_complement cmd/number_complement/main.go

keyboard_row:
	$(GO) build -o bin/keyboard_row cmd/keyboard_row/main.go

test:
	# testing...
	$(GO) test -race -cover $(TEST_PKGS)

