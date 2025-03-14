// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package solve 

import "github.com/llp-devr/push_swap.go/pkg/pushswap" 

func SortThree(ps *pushswap.Stacks) {
	pos := MaxPos(ps.StackA)

	if pos == 0 {
		pushswap.RA(ps)
	} else if pos == 1 {
		pushswap.RRA(ps)
	}

	if (ps.StackA[0] > ps.StackA[1]) {
		pushswap.SA(ps)
	}
}
