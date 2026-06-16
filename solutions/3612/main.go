package main

import (
	"fmt"
	"leetcode-go/lib"
)

func main() {
	fmt.Println(processStr(
		lib.JSONParse[string](`"a#b%*"`),
	))
}