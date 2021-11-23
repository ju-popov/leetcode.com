package main

import (
	"fmt"
	"math"
)

/*

78. Subsets

https://leetcode.com/problems/subsets/

#array #backtracking #bit-manipulation

*/

func subsets(nums []int) [][]int {
	results := make([][]int, 0)
	length := float64(len(nums))

	for i := 0; i < int(math.Pow(2, length)); i++ {
		result := make([]int, 0)

		for j, bitmask := 0, 1; j < len(nums); j++ {
			if (i & bitmask) != 0 {
				result = append(result, nums[j])
			}

			bitmask <<= 1
		}

		results = append(results, result)
	}

	return results
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) // [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
	fmt.Println(subsets([]int{0}))       // [[] [0]]
}
