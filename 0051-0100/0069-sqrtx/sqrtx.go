package main

import "fmt"

/*

69. Sqrt(x)

https://leetcode.com/problems/sqrtx/

*/

func mySqrt(x int) int {
	if (x == 0) || (x == 1) {
		return x
	}

	left := 0
	right := x

	for right-left > 1 {
		mid := left + (right-left)/2

		value := mid * mid

		if value == x {
			return mid
		}

		if mid*mid > x {
			right = mid
		} else {
			left = mid
		}
	}

	return left
}

func main() {
	fmt.Println(mySqrt(0))   // 0
	fmt.Println(mySqrt(1))   // 1
	fmt.Println(mySqrt(2))   // 1
	fmt.Println(mySqrt(3))   // 1
	fmt.Println(mySqrt(4))   // 2
	fmt.Println(mySqrt(5))   // 2
	fmt.Println(mySqrt(6))   // 2
	fmt.Println(mySqrt(7))   // 2
	fmt.Println(mySqrt(8))   // 2
	fmt.Println(mySqrt(9))   // 3
	fmt.Println(mySqrt(10))  // 3
	fmt.Println(mySqrt(255)) // 15
	fmt.Println(mySqrt(256)) // 16
	fmt.Println(mySqrt(257)) // 16
}
