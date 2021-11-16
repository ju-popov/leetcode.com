package main

/*

3. Longest Substring Without Repeating Characters

https://leetcode.com/problems/longest-substring-without-repeating-characters/

Approach 2: Sliding Window

Complexity Analysis

Time complexity : O(2n) = O(n). In the worst case each character will be visited
twice by left and right.

Space complexity : O(min(m, n)). Same as the previous approach. We need O(k)
space for the sliding window, where k is the size of the Map. The size of the
Map is upper bounded by the size of the string n and the size of the
charset/alphabet n.

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
