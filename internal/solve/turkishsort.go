// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package solve

import (
	"github.com/llp-devr/push_swap.go/pkg/pushswap"
)

func TurkishSort(ps *pushswap.Stacks) {
	len_a := len(ps.StackA)
	for i := 0; i < (len_a - 3); i++ {
		pushswap.PB(ps)
	}

	SortThree(ps)

	for 0 < len(ps.StackB) {
		cheapestIdx := findCheapestIdx(ps)
		turkishPA(ps, cheapestIdx)
	}

	minIdx := findMinIdx(ps.StackA)

	if minIdx <= (len(ps.StackA) - minIdx) {
		for i := 0; i < minIdx; i++ {
			pushswap.RA(ps)
		}
	} else {
		for i := 0; i < (len(ps.StackA) - minIdx); i++ {
			pushswap.RRA(ps)
		}
	}
}

func turkishPA(ps *pushswap.Stacks, idx int) {
	targetIdx := findTargetIdx(ps, idx)
	rotationsA := calculateRotations(ps.StackA, targetIdx)
	rotationsB := calculateRotations(ps.StackB, idx)

	if rotationsA < 0 && rotationsB < 0 {
		for i := 0; i < min(abs(rotationsA), abs(rotationsB)); i++ {
			pushswap.RRR(ps)
		}
		for i := 0; i < abs(rotationsA)-min(abs(rotationsA), abs(rotationsB)); i++ {
			pushswap.RRA(ps)
		}
		for i := 0; i < abs(rotationsB)-min(abs(rotationsA), abs(rotationsB)); i++ {
			pushswap.RRB(ps)
		}
	} else if rotationsA > 0 && rotationsB > 0 {
		for i := 0; i < min(rotationsA, rotationsB); i++ {
			pushswap.RR(ps)
		}
		for i := 0; i < rotationsA-min(rotationsA, rotationsB); i++ {
			pushswap.RA(ps)
		}
		for i := 0; i < rotationsB-min(rotationsA, rotationsB); i++ {
			pushswap.RB(ps)
		}
	} else {
		if rotationsA > 0 {
			for i := 0; i < rotationsA; i++ {
				pushswap.RA(ps)
			}
		} else {
			for i := 0; i < abs(rotationsA); i++ {
				pushswap.RRA(ps)
			}
		}
		if rotationsB > 0 {
			for i := 0; i < rotationsB; i++ {
				pushswap.RB(ps)
			}
		} else {
			for i := 0; i < abs(rotationsB); i++ {
				pushswap.RRB(ps)
			}
		}
	}

	pushswap.PA(ps)
}

func findCheapestIdx(ps *pushswap.Stacks) int {
	cheapestIdx := 0
	cheapestCost := 999999

	for i := 0; i < len(ps.StackB); i++ {
		tempCost := calculateCost(ps, i)
		if tempCost < cheapestCost {
			cheapestIdx = i
			cheapestCost = tempCost
		}
	}

	return cheapestIdx
}

func calculateCost(ps *pushswap.Stacks, idx int) int {
	targetIdx := findTargetIdx(ps, idx)
	rotateA := calculateRotations(ps.StackA, targetIdx)
	rotateB := calculateRotations(ps.StackB, idx)

	return minCost(rotateA, rotateB)
}

func calculateRotations(stack []int, idx int) int {
	if idx <= (len(stack) - idx) {
		return idx
	} else {
		return idx - len(stack)
	}
}

func findTargetIdx(ps *pushswap.Stacks, idx int) int {
	minIdx := findMinIdx(ps.StackA)
	maxIdx := findMaxIdx(ps.StackA)

	if ps.StackB[idx] > ps.StackA[maxIdx] {
		return minIdx
	} else {
		for i := minIdx; i < len(ps.StackA); i++ {
			if ps.StackB[idx] < ps.StackA[i] {
				return i
			}
		}
		for i := 0; i < minIdx; i++ {
			if ps.StackB[idx] < ps.StackA[i] {
				return i
			}
		}
	}
	return 0
}

func findMaxIdx(stack []int) int {
	idx := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] > stack[idx] {
			idx = i
		}
	}
	return idx
}

func findMinIdx(stack []int) int {
	idx := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[idx] {
			idx = i
		}
	}
	return idx
}

func minCost(A int, B int) int {
	if (A < 0 && B < 0) || (A >= 0 && B >= 0) {
		if abs(A) > abs(B) {
			return abs(A)
		}
		return abs(B)
	}
	return abs(A) + abs(B)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
