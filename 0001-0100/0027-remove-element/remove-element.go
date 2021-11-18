package main

import "fmt"

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	result := 0

	for _, num := range nums {
		if num != val {
			nums[result] = num
			result++
		}
	}

	return result
}

func main() {
	input1 := []int{3, 2, 2, 3}
	input2 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement(input1, 3), input1)
	fmt.Println(removeElement(input2, 2), input2)
}
