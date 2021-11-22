package main

import "fmt"

/*

77. Combinations

https://leetcode.com/problems/combinations/

#array #backtracking

*/

func combine(n int, k int) [][]int {
	results := make([][]int, 0)

	if k == 1 {
		for i := 1; i <= n; i++ {
			results = append(results, []int{i})
		}

		return results
	}

	for _, prevResult := range combine(n-1, k-1) {
		for i := prevResult[len(prevResult)-1] + 1; i <= n; i++ {
			result := make([]int, 0)
			result = append(result, prevResult...)
			result = append(result, i)
			results = append(results, result)
		}
	}

	return results
}

func main() {
	fmt.Println(combine(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
	fmt.Println(combine(1, 1)) // [[1]]
}
