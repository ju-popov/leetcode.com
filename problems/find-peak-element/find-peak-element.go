package main

import "fmt"

/*

162. Find Peak Element

https://leetcode.com/problems/find-peak-element/

#array #binary-search

*/

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))          // 2
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) // 5
	fmt.Println(findPeakElement([]int{6, 5, 4, 3, 2, 3, 2})) // 0
}
