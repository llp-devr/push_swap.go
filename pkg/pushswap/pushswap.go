// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

type Stacks struct {
	StackA []int
	StackB []int
}

func NewStacks(stack []int) Stacks {
	return Stacks{
		StackA: stack,
		StackB: []int{},
	}
}
