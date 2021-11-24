package main

import "fmt"

/*

45. Jump Game II

https://leetcode.com/problems/jump-game-ii/

*/

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	pos := 0
	result := 0

	for {
		if nums[pos] == 0 {
			return -1
		}

		totalMaxRich := 0
		nextPos := 0

		for step := 1; step <= nums[pos]; step++ {
			if pos+step == len(nums)-1 {
				return result + 1
			}

			maxRich := nums[pos+step] + step

			if maxRich > totalMaxRich {
				totalMaxRich = maxRich
				nextPos = pos + step
			}
		}

		pos = nextPos
		result++
	}
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
