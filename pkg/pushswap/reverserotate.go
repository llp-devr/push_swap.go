// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

import (
	"fmt"
	"strconv"
	"strings"
)

func RRA(ps *PushSwap) {
	reverseRotate(&ps.stackA)
}

func RRB(ps *PushSwap) {
	reverseRotate(&ps.stackB)
}

func RRR(ps *PushSwap) {
	reverseRotate(&ps.stackA)
	reverseRotate(&ps.stackB)
}

func reverseRotate(stack *[]int) {
	if len(*stack) >= 2 {
		*stack = append([]int{(*stack)[len(*stack)-1]}, (*stack)[:len(*stack)-1]...)
	}
}
