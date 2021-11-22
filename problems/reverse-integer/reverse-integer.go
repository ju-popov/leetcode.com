package main

/*

7. Reverse Integer

https://leetcode.com/problems/reverse-integer/

#math

*/

import "fmt"

const (
	maxInt32 = 2147483647
	minInt32 = -2147483648
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	result := int32(0)

	for x != 0 {
		digit := int32(x % 10)

		if result > maxInt32/10 {
			return 0
		}

		if (result == maxInt32/10) && (digit > maxInt32%10) {
			return 0
		}

		if result < minInt32/10 {
			return 0
		}

		if (result == minInt32/10) && (digit < minInt32%10) {
			return 0
		}

		result = result*10 + digit
		x /= 10
	}

	return int(result)
}

func main() {
	fmt.Println(reverse(123))         // 321
	fmt.Println(reverse(-123))        // -321
	fmt.Println(reverse(120))         // 21
	fmt.Println(reverse(0))           // 0
	fmt.Println(reverse(7463847412))  // 2147483647
	fmt.Println(reverse(8463847412))  // 0
	fmt.Println(reverse(-8463847412)) // 2147483648
	fmt.Println(reverse(-9463847412)) // 0
}
