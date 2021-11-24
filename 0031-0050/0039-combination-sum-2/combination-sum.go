package main

import (
	"fmt"
)

/*

39. Combination Sum

https://leetcode.com/problems/combination-sum/

*/

func combinationSumHelper(candidates []int, target int, candidateIndex int, current []int) [][]int {
	results := make([][]int, 0)

	for index := candidateIndex; index < len(candidates); index++ {
		candidate := candidates[index]

		if candidate > target {
			continue
		}

		newCurrent := append([]int{}, current...)
		newCurrent = append(newCurrent, candidate)

		if candidate == target {
			results = append(results, newCurrent)

			continue
		}

		moreSolutions := combinationSumHelper(candidates, target-candidate, index, newCurrent)

		results = append(results, moreSolutions...)
	}

	return results
}

func combinationSum(candidates []int, target int) [][]int {
	return combinationSumHelper(candidates, target, 0, []int{})
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // [[2 2 2 2] [2 3 3] [3 5]]
	fmt.Println(combinationSum([]int{2}, 1))          // []
	fmt.Println(combinationSum([]int{1}, 1))          // [[1]]
	fmt.Println(combinationSum([]int{1}, 2))          // [[1 1]]
}
