package main

import "fmt"

/*

80. Remove Duplicates from Sorted Array II

https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

#array #two-pointers

*/

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	lastNumber := nums[0]
	lastNumberCount := 1

	left := 1
	right := len(nums) - 1
	index := 1

	for index <= right {
		num := nums[index]

		if num != lastNumber {
			lastNumber = num
			lastNumberCount = 1

			if index != left {
				nums[left], nums[index] = nums[index], nums[left]
			}

			left++
			index++

			continue
		}

		if lastNumberCount < 2 {
			lastNumberCount++

			if index != left {
				nums[left], nums[index] = nums[index], nums[left]
			}

			left++
			index++

			continue
		}

		index++
	}

	return left
}

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums1), nums1) // 5 [1 1 2 2 3 1]

	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums2), nums2) // 7 [0 0 1 1 2 3 3 1 1]

	nums3 := []int{1, 2, 2}
	fmt.Println(removeDuplicates(nums3), nums3) // 3 [1 2 2]
}
