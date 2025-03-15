// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

import "fmt"

func SA(ps *Stacks) {
	fmt.Println("sa")

	swap(&ps.StackA)
}

func SB(ps *Stacks) {
	fmt.Println("sb")

	swap(&ps.StackB)
}

func SS(ps *Stacks) {
	fmt.Println("ss")

	swap(&ps.StackA)
	swap(&ps.StackB)
}

func swap(stack *[]int) {
	if len(*stack) >= 2 {
		(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
	}
}
