// https://leetcode.cn/problems/partition-array-according-to-given-pivot/

package main

import (
	"fmt"
	"leetcode-go/utils"
)

func main() {
	// input := "[9,12,5,10,14,3,10]"
	// pivot := 10
	input := "[-3,4,3,2]"
	pivot := 2
	nums := utils.JSONParse[[]int](input)
	ans := pivotArray(*nums, pivot)
	fmt.Println(ans)
}

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, 0, len(nums))
	for _, num := range nums {
		if num < pivot {
			ans = append(ans, num)
		}
	}
	for _, num := range nums {
		if num == pivot {
			ans = append(ans, num)
		}
	}
	for _, num := range nums {
		if num > pivot {
			ans = append(ans, num)
		}
	}
	return ans
}
