// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package solve 

import "github.com/llp-devr/push_swap.go/pkg/pushswap" 

func RadixSort(ps *pushswap.Stacks) {
	bits := bitLenght(Max(ps.StackA))

	for bit := 0; bit < bits; bit++ {
		len_b := len(ps.StackB)
		for i := 0; i < len_b; i++ {
			num := ps.StackB[0]
			if ((num >> bit) & 1) == 1 {
				pushswap.PA(ps)
			} else {
				pushswap.RB(ps)
			}
		}

		len_a := len(ps.StackA)
		for i := 0; i < len_a; i++ {
			num := ps.StackA[0]
			if ((num >> bit) & 1) == 1 {
				pushswap.RA(ps)
			} else {
				pushswap.PB(ps)
			}
		}
	}

	len_b := len(ps.StackB)
	for i := 0; i < len_b; i++ {
		pushswap.PA(ps)
	}
}

func bitLenght(number int) int {
	lenght := 0
	for number > 0 {
	  number >>= 1
		lenght += 1
	}
	return lenght
}
