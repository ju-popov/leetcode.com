package main

import "fmt"

/*

278. First Bad Version

https://leetcode.com/problems/first-bad-version/

#binary-search #interactive

*/

var badVersion int

func isBadVersion(m int) bool {
	return m >= badVersion
}

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	badVersion = 4

	fmt.Println(firstBadVersion(5)) // 4

	badVersion = 1

	fmt.Println(firstBadVersion(1)) // 1

	badVersion = 1

	fmt.Println(firstBadVersion(2)) // 1
}
