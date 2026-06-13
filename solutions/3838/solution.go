package main

import "strings"

//go:generate ../../testgen ./solution.go

var reversed = "zyxwvutsrqponmlkjihgfedcba"

// @input []string ["abcd","def","xyz"]
// @input []int [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
// @want string "rij"
func mapWordWeights(words []string, weights []int) string {
	var b strings.Builder
	for _, word := range words {
		w := 0
		for _, c := range word {
			w += weights[c-'a']
		}
		c := reversed[w%26]
		b.WriteByte(c)
	}
	return b.String()
}
