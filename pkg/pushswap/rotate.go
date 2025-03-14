// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

import "fmt"

func RA(ps *Stacks) {
	fmt.Println("ra")

	rotate(&ps.StackA)
}

func RB(ps *Stacks) {
	fmt.Println("rb")

	rotate(&ps.StackB)
}

func RR(ps *Stacks) {
	fmt.Println("rr")

	rotate(&ps.StackA)
	rotate(&ps.StackB)
}

func rotate(stack *[]int) {
	if len(*stack) >= 2 {
		(*stack) = append((*stack)[1:], (*stack)[0])
	}
}
