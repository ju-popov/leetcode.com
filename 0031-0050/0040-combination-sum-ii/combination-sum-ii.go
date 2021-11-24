package main

import (
	"fmt"
	"sort"
)

/*

40. Combination Sum II

https://leetcode.com/problems/combination-sum-ii/

*/

func combinationSum2Helper(candidates []int, target int, candidateIndex int) [][]int {
	result := make([][]int, 0)

	if candidateIndex >= len(candidates) {
		return result
	}

	candidate := candidates[candidateIndex]

	if candidate == target {
		result = append(result, []int{candidate})
	}

	// With current candidate
	if target > candidate {
		for _, subSolution := range combinationSum2Helper(candidates, target-candidate, candidateIndex+1) {
			solution := append([]int{candidate}, subSolution...)
			result = append(result, solution)
		}
	}

	// Without current candidate - Skip diplicate candidates
	for index := candidateIndex + 1; index < len(candidates); index++ {
		if candidates[index] != candidates[index-1] {
			result = append(result, combinationSum2Helper(candidates, target, index)...)

			break
		}
	}

	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return combinationSum2Helper(candidates, target, 0)
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)) // [[1 1 6] [1 2 5] [1 7] [2 6]]
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))        // [[1 2 2] [5]]
}
