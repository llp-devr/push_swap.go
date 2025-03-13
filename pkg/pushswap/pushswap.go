// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package pushswap

type PushSwap struct {
	stackA []int
	stackB []int
}

func NewPushSwap(stackA []int) *PushSwap {
	return &PushSwap{
		stackA: stackA,
		stackB: []int{},
	}
}
