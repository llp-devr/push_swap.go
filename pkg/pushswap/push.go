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

func PA(ps *PushSwap) {
	if len(ps.stackB) > 0 {
		ps.stackA = append([]int{ps.stackB[0]}, ps.stackA...)
		ps.stackB = ps.stackB[1:]
	}
}

func PA(ps *PushSwap) {
	if len(ps.stackA) > 0 {
		ps.stackB = append([]int{ps.stackA[0]}, ps.stackB...)
		ps.stackA = ps.stackA[1:]
	}
}
