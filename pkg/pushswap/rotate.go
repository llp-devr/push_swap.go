// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

func RA(ps *PushSwap) {
	rotate(&ps.stackA)
}

func RB(ps *PushSwap) {
	rotate(&ps.stackB)
}

func RR(ps *PushSwap) {
	rotate(&ps.stackA)
	rotate(&ps.stackB)
}

func rotate(stack *[]int) {
	if len(*stack) >= 2 {
		*stack = append((*stack)[1:], (*stack)[0])
	}
}
