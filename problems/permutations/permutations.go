package main

import "fmt"

/*

46. Permutations

https://leetcode.com/problems/permutations/

#array #backtracking

*/

func permute(nums []int) [][]int {
	results := make([][]int, 0)

	if len(nums) == 0 {
		results = append(results, []int{})

		return results
	}

	for index := 0; index < len(nums); index++ {
		removedValue := nums[index]

		newNums := append([]int{}, nums[:index]...)
		newNums = append(newNums, nums[index+1:]...)

		subResults := permute(newNums)

		for _, subResult := range subResults {
			result := append([]int{removedValue}, subResult...)
			results = append(results, result)
		}
	}

	return results
}

func main() {
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
	fmt.Println(permute([]int{0, 1}))    // [[0 1] [1 0]]
	fmt.Println(permute([]int{1}))       // [[1]]
	fmt.Println(permute([]int{}))        // [[]]
}
