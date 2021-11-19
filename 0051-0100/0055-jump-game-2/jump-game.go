package main

import "fmt"

/*

55. Jump Game

https://leetcode.com/problems/jump-game/

*/

func canJump(nums []int) bool {
	lastPos := len(nums) - 1

	for index := len(nums) - 1; index >= 0; index-- {
		if index+nums[index] >= lastPos {
			lastPos = index
		}
	}

	return lastPos == 0
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	fmt.Println(canJump([]int{0, 1}))          // false
}
