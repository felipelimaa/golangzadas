package leetcode

import "fmt"

// Write a function that reverses string. The input string is given as an array of characteres s.
// You must do this by modifying the input array in-place with 0(1) extra memory.
func ReverseString(s []string) {
	// Example 1:
	// Input: s["h", "e", "l", "l", "o"]
	// Output: s["o", "l", "l", "e", "h"]

	// Example 2:
	// Input: s["b", "a", "n", "a", "n", "a"]
	// Output: s["a", "n", "a", "n", "a", "b"]

	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l, r = l+1, r-1
	}

	fmt.Println(s)
}
