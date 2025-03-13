// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"

	"github.com/llp-devr/push_swap.go/pkg/converter"
	"github.com/llp-devr/push_swap.go/pkg/pushswap" 
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Error: Please provide exactly one argument.")
		os.Exit(1)
	}

	stack, err := converter.ToIntArray(args[0])
	if err != nil {
		os.Exit(1)
	}

	stacks := pushswap.NewStacks(stack)

	pushswap.PA(&stacks)

	fmt.Println("Message:", stacks)
}
