package main

import (
	"fmt"
)

/*

43. Multiply Strings

https://leetcode.com/problems/multiply-strings/

*/

func addToResult(result []byte, value byte, offset int) {
	for index := len(result) - 1 - offset; index >= 0; index-- {
		if value == 0 {
			return
		}

		total := result[index] - '0' + value

		result[index] = total%10 + '0'

		value = total / 10
	}
}

func multiply(num1 string, num2 string) string {
	// Create result buffer with zeroes
	result := make([]byte, len(num1)+len(num2))
	for index := 0; index < len(result); index++ {
		result[index] = '0'
	}

	// Multiply numbers
	for index2 := len(num2) - 1; index2 >= 0; index2-- {
		for index1 := len(num1) - 1; index1 >= 0; index1-- {
			n1 := num1[index1] - '0'
			n2 := num2[index2] - '0'
			offset := (len(num2) - 1 - index2) + (len(num1) - 1 - index1)

			addToResult(result, n1*n2, offset)
		}
	}

	// Strip leading zeroes
	var index int
	for index = 0; index < len(result)-1; index++ {
		if result[index] != '0' {
			break
		}
	}

	return string(result[index:])
}

func main() {
	fmt.Println(multiply("2", "3"))         // 6
	fmt.Println(multiply("123", "456"))     // 56088
	fmt.Println(multiply("9", "9"))         // 81
	fmt.Println(multiply("9133", "0"))      // 0
	fmt.Println(multiply("0", "52"))        // 0
	fmt.Println(multiply("15999", "15999")) // 255968001
}
