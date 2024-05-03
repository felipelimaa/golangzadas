package main

import (
	"fmt"
	"golangzadas/leetcode"
)

func main() {
	fmt.Println("Welcome to golangzadas")

	var s = []string{"h", "e", "l", "l", "o"}
	leetcode.ReverseString(s)

	word1, word2 := "abc", "pqrs"
	leetcode.MergeStringAlternately(word1, word2)
}
