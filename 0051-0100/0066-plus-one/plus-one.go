package main

import "fmt"

/*

66. Plus One

https://leetcode.com/problems/plus-one/

*/

func plusOne(digits []int) []int {
	result := make([]int, 0)
	addend := []int{1}
	carry := 0

	indexDigits := len(digits) - 1
	indexAddend := len(addend) - 1

	for (indexDigits >= 0) || (indexAddend >= 0) || (carry > 0) {
		sum := carry

		if indexDigits >= 0 {
			sum += digits[indexDigits]
		}

		if indexAddend >= 0 {
			sum += addend[indexAddend]
		}

		carry = sum / 10

		result = append([]int{sum % 10}, result...)

		indexDigits--
		indexAddend--
	}

	return result
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4 3 2 2]
	fmt.Println(plusOne([]int{0}))          // [1]
	fmt.Println(plusOne([]int{9}))          // [1 0]
}
