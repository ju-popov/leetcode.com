package main

import "fmt"

/*

35. Search Insert Position

https://leetcode.com/problems/search-insert-position/

Approach 1: Binary Search

Complexity Analysis

Time complexity : O(logN).

Space complexity : O(1) since it's a constant space solution.

*/

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) // 0
	fmt.Println(searchInsert([]int{1}, 0))          // 0
}
