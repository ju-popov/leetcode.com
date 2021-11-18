package main

/*

13. Roman to Integer

https://leetcode.com/problems/roman-to-integer/

Approach 2: Left-to-Right Pass Improved

Complexity Analysis

Time complexity : O(1).

Space complexity : O(1).

*/

import "fmt"

func romanToInt(s string) int {
	result := 0
	length := len(s)
	left := 0

	for left < length {
		if s[left] == 'M' {
			result += 1000
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'C') && (s[left+1] == 'M') {
			result += 900
			left += 2

			continue
		}

		if s[left] == 'D' {
			result += 500
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'C') && (s[left+1] == 'D') {
			result += 400
			left += 2

			continue
		}

		if s[left] == 'C' {
			result += 100
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'X') && (s[left+1] == 'C') {
			result += 90
			left += 2

			continue
		}

		if s[left] == 'L' {
			result += 50
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'X') && (s[left+1] == 'L') {
			result += 40
			left += 2

			continue
		}

		if s[left] == 'X' {
			result += 10
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'I') && (s[left+1] == 'X') {
			result += 9
			left += 2

			continue
		}

		if s[left] == 'V' {
			result += 5
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'I') && (s[left+1] == 'V') {
			result += 4
			left += 2

			continue
		}

		if s[left] == 'I' {
			result++
			left++

			continue
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("IV"))      // 4
	fmt.Println(romanToInt("IX"))      // 9
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}
