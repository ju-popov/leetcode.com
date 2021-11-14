package main

import "fmt"

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for index, value := range nums {
		if index2, ok := visited[target-value]; ok {
			return []int{index2, index}
		}

		visited[value] = index
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
