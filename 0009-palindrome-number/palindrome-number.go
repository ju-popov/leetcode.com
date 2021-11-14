package main

import "fmt"

func isPalindrome(x int) bool {
	fmt.Println(x)

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
		x = x / 10
	}

	return (x == palindrome) || (palindrome/10 == x)
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(123321))
}
