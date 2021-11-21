package main

import "fmt"

/*

704. Binary Search

https://leetcode.com/problems/binary-search/

#array #binary-search

*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		value := nums[mid]

		if value == target {
			return mid
		}

		if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
}
