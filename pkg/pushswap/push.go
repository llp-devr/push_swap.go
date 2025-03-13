// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

import "fmt"

func PA(ps *Stacks) {
	fmt.Println("pa")

	if len(ps.stackB) > 0 {
		ps.stackA = append([]int{ps.stackB[0]}, ps.stackA...)
		ps.stackB = ps.stackB[1:]
	}
}

func PB(ps *Stacks) {
	fmt.Println("pb")
	if len(ps.stackA) > 0 {
		ps.stackB = append([]int{ps.stackA[0]}, ps.stackB...)
		ps.stackA = ps.stackA[1:]
	}
}
