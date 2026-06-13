package main

import (
	"fmt"
	"slices"
)

//go:generate ../../testgen ./solution.go

// @input []int [1,2]
// @input int 3
// @want int 3
func maxTotalValue(nums []int, k int) int64 {
	maxVal := slices.Max(nums)
	minVal := slices.Min(nums)
	return int64(k * (maxVal - minVal))
}

// 测试
func helper() {
	fmt.Println("hello world")
}
