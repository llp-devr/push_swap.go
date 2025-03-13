# Copyright (c) 2025 Leonardo Lopes Pereira
#
# This file is part of push_swap.go
#
# SPDX-License-Identifier: MIT

BIN ?= push_swap

.PHONY: all
all: push_swap

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f ./push_swap

.PHONY: push_swap
push_swap:
	@echo "Building the project..."
	go build -o push_swap ./cmd/pushswap

.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	go fmt ./...
