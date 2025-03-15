// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"strings"

	"github.com/llp-devr/push_swap.go/pkg/converter"
	"github.com/llp-devr/push_swap.go/pkg/pushswap" 
	"github.com/llp-devr/push_swap.go/internal/solve" 
)

func main() {
	var err error
	var stack []int

	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(1)
	}

	if len(args) == 1 {
		nbrs := strings.Fields(args[0])
		stack, err = converter.ToIntArray(nbrs)
	} else {
		stack, err = converter.ToIntArray(args)
	}
	if err != nil {
		os.Exit(1)
	}

	stacks := pushswap.NewStacks(stack)

	solve.Solve(&stacks)
}
