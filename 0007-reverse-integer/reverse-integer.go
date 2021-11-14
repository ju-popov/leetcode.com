package main

import "fmt"

const minInt32 = -2147483648
const maxInt32 = 2147483647

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	result := int32(0)

	for x != 0 {
		digit := int32(x % 10)

		if x > 0 {
			if result > maxInt32/10 {
				return 0
			}
			if (result == maxInt32/10) && (digit > maxInt32%10) {
				return 0
			}
		} else {
			if result < minInt32/10 {
				return 0
			}
			if (result == minInt32/10) && (digit < minInt32%10) {
				return 0
			}
		}

		result = result*10 + digit
		x = x / 10
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
