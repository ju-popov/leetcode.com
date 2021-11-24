package main

import "fmt"

/*

33. Search in Rotated Sorted Array

https://leetcode.com/problems/search-in-rotated-sorted-array/

#array #binary-search

*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			// left to mid is increasing
			// mid to right is random
			if (nums[left] <= target) && (target < nums[mid]) {
				// target between left and mid (move left)
				right = mid - 1
			} else {
				// move right
				left = mid + 1
			}
		} else {
			// left to mid is random
			// mid to right is increasing
			if (nums[mid] < target) && (target <= nums[right]) {
				// target between mid and right (move right)
				left = mid + 1
			} else {
				// move left
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                   // -1

	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 1)) // 0
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 2)) // 1
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 3)) // 2
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 4)) // 3
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 5)) // -1
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 6)) // 4
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 7)) // 5
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7, 8}, 8)) // 6

	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 1)) // 3
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 2)) // 4
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 3)) // 5
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 4)) // 6
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 5)) // -1
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 6)) // 0
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 7)) // 1
	fmt.Println(search([]int{6, 7, 8, 1, 2, 3, 4}, 8)) // 2
}
