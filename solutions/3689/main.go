// https://leetcode.com/problems/maximum-total-subarray-value-i/

package main

import (
	"fmt"
	"leetcode-go/utils"
	"slices"
)

func main() {
	_input := "[1,3,2]"
	input := utils.JSONParse[[]int](_input)
	k := 2
	ans := maxTotalValue(*input, k)
	fmt.Println(ans)
}

func maxTotalValue(nums []int, k int) int64 {
	maxVal := slices.Max(nums)
	minVal := slices.Min(nums)
	return int64(k * (maxVal - minVal))
}
