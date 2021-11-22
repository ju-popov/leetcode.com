package main

/*

14. Longest Common Prefix

https://leetcode.com/problems/longest-common-prefix/

#string

*/

import "fmt"

func minInt(values ...int) int {
	result := values[0]

	for index := 1; index < len(values); index++ {
		if values[index] < result {
			result = values[index]
		}
	}

	return result
}

func longestCommonPrefixPosition(str1 string, str2 string) int {
	result := -1

	for index := 0; index < minInt(len(str1), len(str2)); index++ {
		if str1[index] == str2[index] {
			result = index

			continue
		}

		break
	}

	return result
}

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	for index := 1; index < len(strs); index++ {
		pos := longestCommonPrefixPosition(result, strs[index])
		if pos < 0 {
			return ""
		}

		result = result[:pos+1]
	}

	return result
}

func main() {
	fmt.Println(longestCommonPrefix(
		[]string{"flower", "flow", "flight"},
	)) // fl
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"})) //
}
