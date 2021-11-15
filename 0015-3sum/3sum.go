package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	if len(nums) < 3 {
		return result
	}

	for index := 0; (index < len(nums)-2) && (nums[index] <= 0); index++ {
		if (index != 0) && (nums[index] == nums[index-1]) {
			continue
		}

		left := index + 1
		right := len(nums) - 1
		targetSum := -nums[index]

		for (left < right) && (nums[left] <= targetSum) {
			if (left != index+1) && (nums[left] == nums[left-1]) {
				left++
				continue
			}

			sum := nums[left] + nums[right]
			if sum > targetSum {
				right--
				continue
			}

			if sum == targetSum {
				result = append(result, []int{nums[index], nums[left], nums[right]})
			}

			left++
		}
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
