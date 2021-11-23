package main

import "fmt"

/*

153. Find Minimum in Rotated Sorted Array

https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

#array #binary-search

*/

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}

		mid := left + (right-left)/2

		if nums[mid] >= nums[0] {
			// left is increasing
			// right is mixed
			// move right
			left = mid + 1
		} else {
			// left is mixed
			// right is increasing
			// move left
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
}
