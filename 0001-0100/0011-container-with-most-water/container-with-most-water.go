package main

/*

11. Container With Most Water

https://leetcode.com/problems/container-with-most-water/

Approach 2: Two Pointer Approach

Complexity Analysis

Time complexity : O(n). Single pass.

Space complexity : O(1). Constant space is used.

*/

import "fmt"

func minValue(values ...int) int {
	result := values[0]

	for index := 1; index < len(values); index++ {
		if values[index] < result {
			result = values[index]
		}
	}

	return result
}

func maxArea(height []int) int {
	maxArea := 0

	leftPos := 0
	rightPos := len(height) - 1

	for leftPos < rightPos {
		leftValue := height[leftPos]
		rightValue := height[rightPos]

		area := minValue(leftValue, rightValue) * (rightPos - leftPos)

		if area > maxArea {
			maxArea = area
		}

		if leftValue <= rightValue {
			leftPos++
		} else {
			rightPos--
		}
	}

	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             // 16
	fmt.Println(maxArea([]int{1, 2, 1}))                   // 2
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))     // 17
}
