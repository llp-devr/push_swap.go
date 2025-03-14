// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package solve 

import "github.com/llp-devr/push_swap.go/pkg/pushswap" 

func Solve(ps *pushswap.Stacks) {
	stackLen := len(ps.StackA)

	if stackLen == 2 {
		SortTwo(ps)
	} else if stackLen == 3 {
		SortThree(ps)
	}
}
