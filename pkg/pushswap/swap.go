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

func SA(ps *PushSwap) {
	swap(&ps.stackA)
}

func SB(ps *PushSwap) {
	swap(&ps.stackB)
}

func SS(ps *PushSwap) {
	swap(&ps.stackA)
	swap(&ps.stackB)
}

func swap(stack *[]int) {
	if len(*stack) >= 2 {
		(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
	}
}
