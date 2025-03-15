// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package solve

func Max(stack []int) int {
	max := stack[0]
	for _, value := range stack {
		if max < value {
			max = value
		}
	}
	return max
}

func MaxPos(stack []int) int {
	pos := 0
	max := stack[0]
	for idx, value := range stack {
		if max < value {
			max = value
			pos = idx
		}
	}
	return pos
}
