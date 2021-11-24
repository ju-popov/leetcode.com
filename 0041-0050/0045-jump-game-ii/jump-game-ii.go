package main

import "fmt"

/*

45. Jump Game II

https://leetcode.com/problems/jump-game-ii/

*/

func jumpHelper(nums []int, pos int, memory map[int]int) int {
	if pos == len(nums)-1 {
		return 0
	}

	memoryKey := pos

	if value, ok := memory[memoryKey]; ok {
		return value
	}

	minResult := -1

	for index := 1; (index <= nums[pos]) && (pos+index < len(nums)); index++ {
		if result := jumpHelper(nums, pos+index, memory); result != -1 {
			if (minResult == -1) || (result < minResult) {
				minResult = result + 1
			}
		}
	}

	memory[memoryKey] = minResult

	return memory[memoryKey]
}

func jump(nums []int) int {
	memory := make(map[int]int)

	return jumpHelper(nums, 0, memory)
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 2
	fmt.Println(jump([]int{1, 2, 3}))       // 2
	fmt.Println(jump([]int{
		5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7,
		9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5,
	})) // 5
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2
	fmt.Println(jump([]int{0}))             // 0
}
