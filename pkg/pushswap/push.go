// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

import "fmt"

func PA(ps *Stacks) {
	fmt.Println("pa")

	if len(ps.StackB) > 0 {
		(*ps).StackA = append([]int{ps.StackB[0]}, ps.StackA...)
		(*ps).StackB = ps.StackB[1:]
	}
}

func PB(ps *Stacks) {
	fmt.Println("pb")

	if len(ps.StackA) > 0 {
		(*ps).StackB = append([]int{ps.StackA[0]}, ps.StackB...)
		(*ps).StackA = ps.StackA[1:]
	}
}
