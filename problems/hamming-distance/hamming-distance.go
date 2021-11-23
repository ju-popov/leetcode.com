package main

import "fmt"

/*

461. Hamming Distance

https://leetcode.com/problems/hamming-distance/

#bit-manipulation

*/

func hammingDistance(x int, y int) int {
	result := 0

	for (x > 0) || (y > 0) {
		if x%2 != y%2 {
			result++
		}

		x /= 2
		y /= 2
	}

	return result
}

func main() {
	fmt.Println(hammingDistance(1, 4)) // 2
	fmt.Println(hammingDistance(3, 1)) // 1
}
