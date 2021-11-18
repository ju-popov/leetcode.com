package main

import (
	"fmt"
	"sort"
)

/*

47. Permutations II

https://leetcode.com/problems/permutations/submissions/

*/

func permuteUniqueHelper(nums []int) [][]int {
	results := make([][]int, 0)

	if len(nums) == 0 {
		results = append(results, []int{})
		return results
	}

	for index := 0; index < len(nums); index++ {
		// Skip duplicate values (in sorted array)
		if index > 0 && (nums[index] == nums[index-1]) {
			continue
		}

		// Remove value at index position
		newNums := append([]int{}, nums[:index]...)
		newNums = append(newNums, nums[index+1:]...)

		for _, subResult := range permuteUniqueHelper(newNums) {
			// value at index position become first
			result := append([]int{nums[index]}, subResult...)
			results = append(results, result)
		}
	}

	return results
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	return permuteUniqueHelper(nums)
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2})) // [[1 1 2] [1 2 1] [2 1 1]]
	fmt.Println(permuteUnique([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
