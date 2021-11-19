package main

import "fmt"

/*

47. Permutations II

https://leetcode.com/problems/permutations/submissions/

*/

func permuteUnique(nums []int) [][]int {
	results := make([][]int, 0)

	if len(nums) == 0 {
		results = append(results, []int{})

		return results
	}

	memory := make(map[int]bool)

	for index := 0; index < len(nums); index++ {
		removedValue := nums[index]

		if memory[removedValue] {
			continue
		}

		memory[removedValue] = true

		newNums := append([]int{}, nums[:index]...)
		newNums = append(newNums, nums[index+1:]...)

		subResults := permuteUnique(newNums)

		for _, subResult := range subResults {
			result := append([]int{removedValue}, subResult...)
			results = append(results, result)
		}
	}

	return results
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2})) // [[1 1 2] [1 2 1] [2 1 1]]
	fmt.Println(permuteUnique([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
