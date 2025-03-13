// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

type Stacks struct {
	stackA []int
	stackB []int
}

func NewStacks(stackA []int) Stacks {
	return Stacks{
		stackA: stackA,
		stackB: []int{},
	}
}
