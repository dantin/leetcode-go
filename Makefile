
CURDIR := $(shell pwd)
GO := go

TEST_PKGS := $(shell find . -iname "*_test.go" -exec dirname {} \;)

.PHONY: build hamming_distance array_partition_i number_complement \
	keyboard_row reverse_words_in_a_string_iii distribute_candies

build: hamming_distance array_partition_i number_complement keyboard_row \
	reverse_words_in_a_string_iii distribute_candies

hamming_distance:
	$(GO) build -o bin/hamming_distance cmd/hamming_distance/main.go

array_partition_i:
	$(GO) build -o bin/array_partition_i cmd/array_partition_i/main.go

number_complement:
	$(GO) build -o bin/number_complement cmd/number_complement/main.go

keyboard_row:
	$(GO) build -o bin/keyboard_row cmd/keyboard_row/main.go

reverse_words_in_a_string_iii:
	$(GO) build -o bin/reverse_words_in_a_string_iii cmd/reverse_words_in_a_string_iii/main.go

distribute_candies:
	$(GO) build -o bin/distribute_candies cmd/distribute_candies/main.go

test:
	# testing...
	$(GO) test -race -cover $(TEST_PKGS)

