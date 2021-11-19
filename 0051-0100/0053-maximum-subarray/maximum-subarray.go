package main

import "fmt"

/*

53. Maximum Subarray

https://leetcode.com/problems/maximum-subarray/

*/

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for index := 1; index < len(nums); index++ {
		if currentSum < 0 {
			currentSum = 0
		}

		currentSum += nums[index]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
}
