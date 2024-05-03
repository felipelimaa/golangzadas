package leetcode

import (
	"fmt"
)

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
// If a string is longer than the other, append the additional letters onto the end of the merged string.
func MergeStringAlternately(word1, word2 string) {
	var r string
	var idx int

	for idx < len(word1) && idx < len(word2) {
		r = r + string(word1[idx]) + string(word2[idx])
		idx++
	}

	if idx < len(word1) {
		r = r + string(word1[idx:])
		idx++
	}

	if idx < len(word2) {
		r = r + string(word2[idx:])
		idx++
	}

	fmt.Println(r)
}
