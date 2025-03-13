// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package converter

import (
	"fmt"
	"strconv"
	"strings"
)

func ToIntArray(input string) ([]int, error) {
	strs := strings.Fields(input)

	var stack []int

	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to an integer: %v", str, err)
		}
		stack = append(stack, num)
	}

	return stack, nil
}
