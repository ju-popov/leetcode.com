package main

import "fmt"

/*

81. Search in Rotated Sorted Array II

https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

*/

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return true
		}

		if target == nums[left] {
			return true
		}

		if target == nums[right] {
			return true
		}

		if nums[mid] >= nums[left] {
			// left ordered
			// right random
			if (nums[left] < target) && (target < nums[mid]) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// left random
			// right ordered
			if (nums[mid] < target) && (target < nums[right]) {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))       // true
}
