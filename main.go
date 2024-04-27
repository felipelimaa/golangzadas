package main

import (
	"fmt"
	"golangzadas/leetcode"
)

func main() {
	fmt.Println("Welcome to golangzadas")

	// Leetcode
	// Challenge 344
	var s []string
	s = []string{"h", "e", "l", "l", "o"}
	leetcode.ReverseString(s)
	s = []string{"b", "a", "n", "a", "n", "a"}
	leetcode.ReverseString(s)

	word1, word2 := "abc", "pqrs"
	leetcode.MergeStringAlternately(word1, word2)
}
