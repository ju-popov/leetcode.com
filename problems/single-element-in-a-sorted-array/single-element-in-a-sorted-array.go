package main

import "fmt"

/*

540. Single Element in a Sorted Array

https://leetcode.com/problems/single-element-in-a-sorted-array/

#array #binary-search

*/

func singleNonDuplicate(nums []int) int {
	left := -1
	right := len(nums) / 2

	for {
		if right-left == 1 {
			return nums[right*2]
		}

		mid := left + (right-left)/2

		if nums[mid*2] == nums[mid*2+1] {
			left = mid
		} else {
			right = mid
		}
	}
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1}))                         // 1
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})) // 2
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))    // 10
}
