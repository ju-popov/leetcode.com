package main

import "fmt"

/*

55. Jump Game

https://leetcode.com/problems/jump-game/

*/

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	pos := 0

	for {
		nextPos := pos
		maxJumpTo := 0

		for step := 1; step <= nums[pos]; step++ {
			jumpTo := pos + step + nums[pos+step]
			if jumpTo >= len(nums)-1 {
				return true
			}

			if jumpTo > maxJumpTo {
				maxJumpTo = jumpTo
				nextPos = pos + step
			}
		}

		if nextPos == pos {
			return false
		}

		pos = nextPos
	}
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	fmt.Println(canJump([]int{0, 1}))          // false
}
