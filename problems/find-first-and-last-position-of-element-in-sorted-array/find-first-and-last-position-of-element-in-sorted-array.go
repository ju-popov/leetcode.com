package main

import (
	"fmt"
)

/*

34. Find First and Last Position of Element in Sorted Array

https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

#array #binary-search

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
