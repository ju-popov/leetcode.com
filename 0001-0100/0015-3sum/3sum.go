package main

/*

15. 3Sum

https://leetcode.com/problems/3sum/

Approach 1: Two Pointers

Complexity Analysis

Time Complexity: {O}(n^2).

Sorting the array takes {O}(nlog{n}), so overall complexity is
{O}(nlog{n} + n^2). This is asymptotically equivalent to {O}(n^2).

Space Complexity: from {O}(log{n}) {O}(n), depending on the implementation of
the sorting algorithm. For the purpose of complexity analysis, we ignore the
memory required for the output.

*/

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	target := 0
	result := make([][]int, 0)

	if len(nums) < 3 {
		return result
	}

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

			sum := nums[index] + nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[index], nums[left], nums[right]})
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
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{}))                    // []
	fmt.Println(threeSum([]int{0}))                   // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [0, 0, 0]
	fmt.Println(threeSum([]int{0, 0, 0, 0}))          // [0, 0, 0]
}
