// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

func RA(ps *Stacks) {
	rotate(&ps.stackA)
}

func RB(ps *Stacks) {
	rotate(&ps.stackB)
}

func RR(ps *Stacks) {
	rotate(&ps.stackA)
	rotate(&ps.stackB)
}

func rotate(stack *[]int) {
	if len(*stack) >= 2 {
		*stack = append((*stack)[1:], (*stack)[0])
	}
}
