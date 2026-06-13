package main

import (
	"fmt"
	"leetcode-go/lib"
)

func main() {
	fmt.Println(maxTotalValue(
		lib.JSONParse[[]int](`[1,2]`),
		lib.JSONParse[int](`3`),
	))
}