// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

import "fmt"

func RRA(ps *Stacks) {
	fmt.Println("rra")
	
	reverseRotate(&ps.StackA)
}

func RRB(ps *Stacks) {
	fmt.Println("rrb")

	reverseRotate(&ps.StackB)
}

func RRR(ps *Stacks) {
	fmt.Println("rrr")
	
	reverseRotate(&ps.StackA)
	reverseRotate(&ps.StackB)
}

func reverseRotate(stack *[]int) {
	if len(*stack) >= 2 {
		(*stack) = append([]int{(*stack)[len(*stack)-1]}, (*stack)[:len(*stack)-1]...)
	}
}
