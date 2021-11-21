package main

import "fmt"

/*

374. Guess Number Higher or Lower

https://leetcode.com/problems/guess-number-higher-or-lower/

#binary-search #iteractive

*/

var pick int

func guess(num int) int {
	if pick == num {
		return 0
	}

	if pick > num {
		return 1
	}

	return -1
}

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := left + (right-left)/2

		switch guess(mid) {
		case 1:
			left = mid + 1
		case -1:
			right = mid - 1
		default:
			return mid
		}
	}

	return -1
}

func main() {
	pick = 6

	fmt.Println(guessNumber(10)) // 6

	pick = 1

	fmt.Println(guessNumber(1)) // 1

	pick = 1

	fmt.Println(guessNumber(2)) // 1

	pick = 2

	fmt.Println(guessNumber(2)) // 2
}
