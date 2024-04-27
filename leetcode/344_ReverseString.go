package leetcode

import "fmt"

// Write a function that reverses string. The input string is given as an array of characteres s.
// You must do this by modifying the input array in-place with 0(1) extra memory.
func ReverseString(s []string) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l, r = l+1, r-1
	}

	fmt.Println(s)
}
