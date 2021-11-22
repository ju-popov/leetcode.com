package main

/*

1. Two Sum

https://leetcode.com/problems/two-sum/

#array #hash-table

*/

import "fmt"

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for index, value := range nums {
		if _, ok := visited[target-value]; ok {
			return []int{visited[target-value], index}
		}

		visited[value] = index
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1, 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0, 1]
}
