package main

/*

8. String to Integer (atoi)

https://leetcode.com/problems/string-to-integer-atoi/

*/

import "fmt"

const (
	maxInt32 = 2147483647
	minInt32 = -2147483648
)

func myAtoi(s string) int {
	left := 0
	length := len(s)

	for (left < length) && (s[left] == ' ') {
		left++
	}

	positive := true

	if left < length {
		if s[left] == '-' {
			positive = false
			left++
		} else if s[left] == '+' {
			left++
		}
	}

	result := int32(0)

	for left < length {
		if s[left] < '0' || s[left] > '9' {
			break
		}

		digit := int32(s[left] - '0')

		if positive {
			if result > maxInt32/10 {
				result = maxInt32

				break
			}

			if (result == maxInt32/10) && (digit > maxInt32%10) {
				result = maxInt32

				break
			}
		} else {
			if result > -minInt32/10 {
				result = minInt32

				break
			}

			if (result == -minInt32/10) && (digit > -minInt32%10) {
				result = minInt32

				break
			}
		}

		result = result*10 + digit
		left++
	}

	if !positive {
		result = -result
	}

	return int(result)
}

func main() {
	fmt.Println(myAtoi("42"))              // 42
	fmt.Println(myAtoi("   -42"))          // -42
	fmt.Println(myAtoi("4193 with words")) // 4193
	fmt.Println(myAtoi("words and 987"))   // 0
	fmt.Println(myAtoi("-91283472332"))    // -2147483648
	fmt.Println(myAtoi("2147483646"))      // 2147483646
	fmt.Println(myAtoi("2147483647"))      // 2147483647
	fmt.Println(myAtoi("2147483648"))      // 2147483647
	fmt.Println(myAtoi("-2147483647"))     // -2147483647
	fmt.Println(myAtoi("-2147483648"))     // -2147483648
	fmt.Println(myAtoi("-2147483649"))     // -2147483648
}
