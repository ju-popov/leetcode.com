package main

import (
	"fmt"
)

/*

39. Combination Sum

https://leetcode.com/problems/combination-sum/

*/

type memoryKey struct {
	candidateIndex int
	target         int
}

func combineSolution(candidate int, quantity int, add []int) []int {
	result := make([]int, 0)

	for index := 0; index < quantity; index++ {
		result = append(result, candidate)
	}

	if add != nil {
		result = append(result, add...)
	}

	return result
}

func combinationSumHelper(candidates []int, candidateIndex int, target int, memory map[memoryKey][][]int) [][]int {
	result := make([][]int, 0)

	if candidateIndex >= len(candidates) {
		return result
	}

	key := memoryKey{candidateIndex: candidateIndex, target: target}
	if value, ok := memory[key]; ok {
		return value
	}

	candidate := candidates[candidateIndex]

	for quantity := target / candidate; quantity >= 0; quantity-- {
		remaining := target - candidate*quantity

		if remaining == 0 {
			result = append(result, combineSolution(candidate, quantity, nil))

			continue
		}

		for _, additional := range combinationSumHelper(candidates, candidateIndex+1, remaining, memory) {
			result = append(result, combineSolution(candidate, quantity, additional))
		}
	}

	memory[key] = result

	return memory[key]
}

func combinationSum(candidates []int, target int) [][]int {
	memory := make(map[memoryKey][][]int)

	return combinationSumHelper(candidates, 0, target, memory)
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // [[2 2 2 2] [2 3 3] [3 5]]
	fmt.Println(combinationSum([]int{2}, 1))          // []
	fmt.Println(combinationSum([]int{1}, 1))          // [[1]]
	fmt.Println(combinationSum([]int{1}, 2))          // [[1 1]]
}
