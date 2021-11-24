package main

import (
	"fmt"
	"sort"
	"strings"
)

/*

49. Group Anagrams

https://leetcode.com/problems/group-anagrams/

#hash-table #string #sorting

*/

func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)

	return strings.Join(chars, "")
}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	memory := make(map[string][]string)

	for _, str := range strs {
		sorted := sortString(str)
		memory[sorted] = append(memory[sorted], str)
	}

	for _, values := range memory {
		result = append(result, values)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams(
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
	)) // [[eat tea ate] [tan nat] [bat]]
	fmt.Println(groupAnagrams([]string{""}))  // [[]]
	fmt.Println(groupAnagrams([]string{"a"})) // [[a]]
}
