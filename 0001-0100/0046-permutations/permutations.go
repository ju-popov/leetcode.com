package main

import "fmt"

/*

46. Permutations

https://leetcode.com/problems/permutations/

*/

func permute(nums []int) [][]int {
	results := make([][]int, 0)

	if len(nums) == 0 {
		results = append(results, []int{})
		return results
	}

	for _, subResult := range permute(nums[1:]) {
		for index := 0; index < len(nums); index++ {
			result := append([]int{}, subResult[:index]...)
			result = append(result, nums[0])
			result = append(result, subResult[index:]...)
			results = append(results, result)
		}
	}

	return results
}

func main() {
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [2 1 3] [2 3 1] [1 3 2] [3 1 2] [3 2 1]]
	fmt.Println(permute([]int{0, 1}))    // [[0 1] [1 0]]
	fmt.Println(permute([]int{1}))       // [[1]]
	fmt.Println(permute([]int{}))        // [[]]
}
