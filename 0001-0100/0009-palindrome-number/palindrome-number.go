package main

/*

9. Palindrome Number

https://leetcode.com/problems/palindrome-number/

Approach 1: Revert half of the number

Complexity Analysis

Time complexity : O(log10(n)). We divided the input by 10 for every iteration,
so the time complexity is O(log10(n))
​
Space complexity : O(1).

*/

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x < 10 {
		return true
	}

	if (x % 10) == 0 {
		return false
	}

	palindrome := 0

	for x > palindrome {
		palindrome = (palindrome * 10) + (x % 10)
		x /= 10
	}

	return (x == palindrome) || (palindrome/10 == x)
}

func main() {
	fmt.Println(isPalindrome(121))    // true
	fmt.Println(isPalindrome(-121))   // false
	fmt.Println(isPalindrome(0))      // true
	fmt.Println(isPalindrome(10))     // false
	fmt.Println(isPalindrome(-101))   // false
	fmt.Println(isPalindrome(1))      // true
	fmt.Println(isPalindrome(11))     // true
	fmt.Println(isPalindrome(121))    // true
	fmt.Println(isPalindrome(1221))   // true
	fmt.Println(isPalindrome(12321))  // true
	fmt.Println(isPalindrome(123321)) // true
}
