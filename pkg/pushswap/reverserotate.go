// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

func RRA(ps *Stacks) {
	reverseRotate(&ps.stackA)
}

func RRB(ps *Stacks) {
	reverseRotate(&ps.stackB)
}

func RRR(ps *Stacks) {
	reverseRotate(&ps.stackA)
	reverseRotate(&ps.stackB)
}

func reverseRotate(stack *[]int) {
	if len(*stack) >= 2 {
		*stack = append([]int{(*stack)[len(*stack)-1]}, (*stack)[:len(*stack)-1]...)
	}
}
