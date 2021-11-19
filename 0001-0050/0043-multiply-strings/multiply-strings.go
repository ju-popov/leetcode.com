package main

import "fmt"

/*

43. Multiply Strings

https://leetcode.com/problems/multiply-strings/

*/

func addition(num1 string, num2 string) string {
	var carry byte

	indexNum1 := len(num1) - 1
	indexNum2 := len(num2) - 1
	result := ""

	for (carry != 0) || (indexNum1 >= 0) || (indexNum2 >= 0) {
		total := carry

		if indexNum1 >= 0 {
			total += num1[indexNum1] - '0'
		}

		if indexNum2 >= 0 {
			total += num2[indexNum2] - '0'
		}

		carry = total / 10

		indexNum1--
		indexNum2--

		result = string(total%10+'0') + result
	}

	return result
}

func multipyByByte(num1 string, num2 byte, rightPadding int) string {
	result := ""

	var carry byte

	indexNum1 := len(num1) - 1

	for (carry != 0) || (indexNum1 >= 0) {
		total := carry

		if indexNum1 >= 0 {
			total += (num1[indexNum1] - '0') * num2
		}

		carry = total / 10

		indexNum1--

		result = string(total%10+'0') + result
	}

	for index := 0; index < rightPadding; index++ {
		result += "0"
	}

	if result[0] == '0' {
		result = "0"
	}

	return result
}

func multiply(num1 string, num2 string) string {
	result := "0"

	for index := len(num2) - 1; index >= 0; index-- {
		result = addition(result, multipyByByte(num1, num2[index]-'0', len(num2)-1-index))
	}

	return result
}

func main() {
	fmt.Println(multiply("2", "3"))         // 6
	fmt.Println(multiply("123", "456"))     // 56088
	fmt.Println(multiply("9", "9"))         // 81
	fmt.Println(multiply("9133", "0"))      // 0
	fmt.Println(multiply("0", "52"))        // 0
	fmt.Println(multiply("15999", "15999")) // 255968001
}
