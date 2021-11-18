package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	result := 0
	lastNumber := 0

	for _, num := range nums {
		if (result == 0) || (num != lastNumber) {
			nums[result] = num
			lastNumber = num
			result++
		}
	}

	return result
}

func main() {
	input1 := []int{1, 1, 2}
	input2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(input1), input1)
	fmt.Println(removeDuplicates(input2), input2)
}
