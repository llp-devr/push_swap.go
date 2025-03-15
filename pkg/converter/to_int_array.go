// Copyright (c) 2025 Leonardo Lopes Pereira
//
// This file is part of push_swap.go
//
// SPDX-License-Identifier: MIT

package converter

import (
	"fmt"
	"sort"
	"strconv"
)

func ToIntArray(strs []string) ([]int, error) {
	var list []int

	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to an integer: %v", str, err)
		}
		list = append(list, num)
	}

	sorted := make([]int, len(list))
	copy(sorted, list)
	sort.Ints(sorted)

	var stack []int
	for _, el := range list {
		stack = append(stack, findIndex(sorted, el))
	}

	return stack, nil
}

func findIndex(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}
