package main

import "fmt"

// https://leetcode.com/problems/implement-strstr/

func isMatch(haystack string, needle string) bool {
	for index := 0; index < len(needle); index++ {
		if haystack[index] != needle[index] {
			return false
		}
	}

	return true
}

func strStr(haystack string, needle string) int {
	for index := 0; index < len(haystack)-len(needle)+1; index++ {
		if isMatch(haystack[index:], needle) {
			return index
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
}
