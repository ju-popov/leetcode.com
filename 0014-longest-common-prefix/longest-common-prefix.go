package main

/*

14. Longest Common Prefix

https://leetcode.com/problems/longest-common-prefix/

Approach 1: Horizontal scanning

Complexity Analysis

Time complexity : O(S) , where S is the sum of all characters in all strings.

In the worst case all n strings are the same. The algorithm compares the string
S1 with the other strings [S2..SN]. There are S character comparisons, where S
is the sum of all characters in the input array.

Space complexity : O(1). We only used constant extra space.

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
