package main

import "fmt"

/*

75. Sort Colors

https://leetcode.com/problems/sort-colors/

#array #two-pointers #sorting

*/

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	index := 0

	for index <= right {
		switch nums[index] {
		case 0:
			nums[left], nums[index] = nums[index], nums[left]
			left++
			index++
		case 2:
			nums[right], nums[index] = nums[index], nums[right]
			right--
		default:
			index++
		}
	}
}

func main() {
	nums1 := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums1)

	fmt.Println(nums1)

	nums2 := []int{2, 0, 1}

	sortColors(nums2)

	fmt.Println(nums2)

	nums3 := []int{0}

	sortColors(nums3)

	fmt.Println(nums3)

	nums4 := []int{1}

	sortColors(nums4)

	fmt.Println(nums4)
}
