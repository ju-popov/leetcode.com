package main

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

func longestCommonPrefixHelper(str1 string, str2 string) int {
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
	str1 := strs[0]

	for index := 1; index < len(strs); index++ {
		result := longestCommonPrefixHelper(str1, strs[index])
		if result < 0 {
			return ""
		}
		str1 = str1[:result+1]
	}

	return str1
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    //
}
