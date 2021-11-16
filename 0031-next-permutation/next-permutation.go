package main

import "fmt"

// https://leetcode.com/problems/next-permutation/solution/

func nextPermutation(nums []int) []int {
	// Swap Left Value
	swapLeftIndex := -1
	for swapLeftIndex = len(nums) - 2; swapLeftIndex >= 0; swapLeftIndex-- {
		if nums[swapLeftIndex] < nums[swapLeftIndex+1] {
			break
		}
	}

	// Swap Right Value
	if swapLeftIndex != -1 {
		swapRightIndex := -1

		for index := len(nums) - 1; index > swapLeftIndex; index-- {
			if nums[index] > nums[swapLeftIndex] {
				if (swapRightIndex == -1) || (nums[index] < nums[swapRightIndex]) {
					swapRightIndex = index
				}
			}
		}

		nums[swapLeftIndex], nums[swapRightIndex] = nums[swapRightIndex], nums[swapLeftIndex]
	}

	// Reverse the rest (sort ascending)
	start := swapLeftIndex + 1
	end := len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	return nums
}

func main() {
	fmt.Println(nextPermutation([]int{1, 2, 3}))
	fmt.Println(nextPermutation([]int{1, 3, 2}))
	fmt.Println(nextPermutation([]int{2, 1, 3}))
	fmt.Println(nextPermutation([]int{2, 3, 1}))
	fmt.Println(nextPermutation([]int{3, 1, 2}))
	fmt.Println(nextPermutation([]int{3, 2, 1}))
	fmt.Println(nextPermutation([]int{2, 3, 1, 3, 3}))
}
