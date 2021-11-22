package main

/*

18. 4Sum

https://leetcode.com/problems/4sum/

#array #two-pointers #sorting

*/

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)

	for index1 := 0; index1 < len(nums)-3; index1++ {
		if (index1 != 0) && (nums[index1] == nums[index1-1]) {
			continue
		}

		for index2 := index1 + 1; index2 < len(nums)-2; index2++ {
			if (index2 != index1+1) && (nums[index2] == nums[index2-1]) {
				continue
			}

			left := index2 + 1
			right := len(nums) - 1

			for left < right {
				if (left != index2+1) && (nums[left] == nums[left-1]) {
					left++

					continue
				}

				sum := nums[index1] + nums[index2] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[index1], nums[index2], nums[left], nums[right]})
				}

				if sum > target {
					right--

					continue
				}

				left++
			}
		}
	}

	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))   // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))        // [[2,2,2,2]]
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, -1)) // [[-4,0,1,2],[-1,-1,0,1]]
}
