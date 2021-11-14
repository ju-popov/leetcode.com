package main

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

func findNextPos(height []int, left int, right int, biggerThan int, direction bool) int {
	if direction {
		for index := left; index <= right; index++ {
			if height[index] > biggerThan {
				return index
			}
		}
	} else {
		for index := right; index >= left; index-- {
			if height[index] > biggerThan {
				return index
			}
		}
	}

	return -1
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
			if result := findNextPos(height, leftPos+1, rightPos-1, leftValue, true); result >= 0 {
				leftPos = result
			} else {
				break
			}
		} else {
			if result := findNextPos(height, leftPos+1, rightPos-1, rightValue, false); result >= 0 {
				rightPos = result
			} else {
				break
			}
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
