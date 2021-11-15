package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, targetSum int) [][]int {
	result := make([][]int, 0)

	left := 0
	right := len(nums) - 1

	prevLeftValue := 0
	for (left < right) && (nums[left] <= targetSum) {
		leftValue := nums[left]

		if (left == 0) || (leftValue != prevLeftValue) {
			rightValue := nums[right]
			sum := leftValue + rightValue

			if sum > targetSum {
				right--
				continue
			}

			if sum == targetSum {
				result = append(result, []int{leftValue, rightValue})
			}
		}

		prevLeftValue = leftValue
		left++
	}

	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	if len(nums) < 3 {
		return result
	}

	prevTargetSum := 0
	for index := 0; index < len(nums)-2; index++ {
		targetSum := -nums[index]
		if (index == 0) || (targetSum != prevTargetSum) {
			results := twoSum(nums[index+1:], targetSum)
			for _, value := range results {
				result = append(result, []int{nums[index], value[0], value[1]})
			}
		}

		prevTargetSum = targetSum
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{}))                    // []
	fmt.Println(threeSum([]int{0}))                   // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [0, 0, 0]
	fmt.Println(threeSum([]int{0, 0, 0, 0}))          // [0, 0, 0]
}
