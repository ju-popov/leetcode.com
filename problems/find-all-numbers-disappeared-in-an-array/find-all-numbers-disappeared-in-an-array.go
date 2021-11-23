package main

import "fmt"

/*

448. Find All Numbers Disappeared in an Array

https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

#array #hash-table

*/

func findDisappearedNumbers(nums []int) []int {
	memory := make([]bool, len(nums))

	for _, num := range nums {
		memory[num-1] = true
	}

	result := make([]int, 0)

	for index := 0; index < len(nums); index++ {
		if !memory[index] {
			result = append(result, index+1)
		}
	}

	return result
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
