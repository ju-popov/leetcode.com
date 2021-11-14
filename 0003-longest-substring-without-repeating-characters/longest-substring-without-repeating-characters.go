package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxResult := 0

	subString := ""

	for _, char := range s {
		if index := strings.Index(subString, string(char)); index >= 0 {
			subString = subString[index+1:]
		}

		subString += string(char)

		if result := len(subString); result > maxResult {
			maxResult = result
		}
	}

	return maxResult
}

func main() {
	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("", lengthOfLongestSubstring(""))
	fmt.Println("dvdf", lengthOfLongestSubstring("dvdf"))
}
