package main

import (
	"fmt"
	"leetcode-go/lib"
)

func main() {
	fmt.Println(mapWordWeights(
		lib.JSONParse[[]string](`["abcd","def","xyz"]`),
		lib.JSONParse[[]int](`[5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]`),
	))
}