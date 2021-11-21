package main

/*

3. Longest Substring Without Repeating Characters

https://leetcode.com/problems/longest-substring-without-repeating-characters/

*/

import (
	"fmt"
)

func maxInt(values ...int) int {
	result := values[0]

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	visited := make(map[byte]bool)

	left := -1
	right := 0

	for right < len(s) {
		char := s[right]

		if visited[char] {
			for {
				left++

				visited[s[left]] = false

				if s[left] == char {
					break
				}
			}
		}

		visited[char] = true

		result = maxInt(result, right-left)

		right++
	}

	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring(""))         // 0
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
}
