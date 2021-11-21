package main

import "fmt"

/*

66. Plus One

https://leetcode.com/problems/plus-one/

*/

func plusOne(digits []int) []int {
	carry := 1
	indexDigits := len(digits) - 1

	for (indexDigits >= 0) && (carry > 0) {
		sum := carry + digits[indexDigits]
		carry = sum / 10
		digits[indexDigits] = sum % 10
		indexDigits--
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4 3 2 2]
	fmt.Println(plusOne([]int{0}))          // [1]
	fmt.Println(plusOne([]int{9}))          // [1 0]
}
