package main

import "fmt"

/*

33. Search in Rotated Sorted Array

https://leetcode.com/problems/search-in-rotated-sorted-array/

Approach 2: One-pass Binary Search

Complexity Analysis

Time complexity : O(logN).

Space complexity : O(1).

*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			// target found
			return mid
		}

		if nums[mid] >= nums[left] {
			// left part if increasing
			// right side is mixed
			if target >= nums[left] && target < nums[mid] {
				// target is in the left increasaing part
				right = mid - 1
			} else {
				// target is in the right mixed part
				left = mid + 1
			}

			continue
		}

		// left part if mixes
		// right side is increasing
		if target > nums[mid] && target <= nums[right] {
			// target is in the right increasing part
			left = mid + 1
		} else {
			// target is in the left mixed part
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                   // -1
}
