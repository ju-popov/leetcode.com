package main

/*

16. 3Sum Closest

https://leetcode.com/problems/3sum-closest/

#array #two-pointers #sorting

*/

import (
	"fmt"
	"sort"
)

func calcResult(n1 int, n2 int, n3 int, target int) (int, int) {
	total := n1 + n2 + n3

	offset := target - total
	if offset < 0 {
		offset = -offset
	}

	return total, offset
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	result, resultOffset := calcResult(nums[0], nums[1], nums[2], target)

	for index := 0; index < len(nums)-2; index++ {
		if (index != 0) && (nums[index] == nums[index-1]) {
			continue
		}

		left := index + 1
		right := len(nums) - 1

		for left < right {
			if (left != index+1) && (nums[left] == nums[left-1]) {
				left++

				continue
			}

			sum, sumOffset := calcResult(nums[index], nums[left], nums[right], target)

			if sumOffset == 0 {
				return sum
			}

			if sumOffset < resultOffset {
				result, resultOffset = sum, sumOffset
			}

			if sum > target {
				right--

				continue
			}

			left++
		}
	}

	return result
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))         // 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))              // 0
	fmt.Println(threeSumClosest([]int{-100, -98, -2, -1}, -101)) // -101
}
