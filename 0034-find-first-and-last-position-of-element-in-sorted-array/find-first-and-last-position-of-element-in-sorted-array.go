package main

import (
	"fmt"
)

/*

34. Find First and Last Position of Element in Sorted Array

https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

Approach: Binary Search

Complexity Analysis

Time Complexity: O(logN) considering there are N elements in the array. This is
because binary search takes logarithmic time to scan an array of N elements.
Why? Because at each step we discard half of the array we are scanning and
hence, we're done after a logarithmic number of steps. We simply perform binary
search twice in this case.

Space Complexity: O(1) since we only use space for a few variables and our
result array, all of which require constant space.

*/

func binarySeach(nums []int, target int, first bool) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			if first {
				if (mid == left) || (nums[mid-1] != target) {
					return mid
				}

				right = mid - 1
			} else {
				if (mid == right) || (nums[mid+1] != target) {
					return mid
				}

				left = mid + 1
			}

			continue
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := binarySeach(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}

	last := binarySeach(nums, target, false)

	return []int{first, last}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3,4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1, -1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1, -1]
}
