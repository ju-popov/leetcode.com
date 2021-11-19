package main

/*

12. Integer to Roman

https://leetcode.com/problems/integer-to-roman/

Approach 1: Greedy

Complexity Analysis

Time complexity : O(1).

As there is a finite set of roman numerals, there is a hard upper limit on how
many times the loop can iterate. This upper limit is 15 times, and it occurs for
the number 3888, which has a representation of MMMDCCCLXXXVIII. Therefore, we
say the time complexity is constant, i.e. O(1).

Space complexity : O(1).

The amount of memory used does not change with the size of the input integer,
and is therefore constant.

*/

import "fmt"

func intToRoman(num int) string {
	result := ""

	for num >= 1000 {
		result += "M"
		num -= 1000
	}

	for num >= 900 {
		result += "CM"
		num -= 900
	}

	for num >= 500 {
		result += "D"
		num -= 500
	}

	for num >= 400 {
		result += "CD"
		num -= 400
	}

	for num >= 100 {
		result += "C"
		num -= 100
	}

	for num >= 90 {
		result += "XC"
		num -= 90
	}

	for num >= 50 {
		result += "L"
		num -= 50
	}

	for num >= 40 {
		result += "XL"
		num -= 40
	}

	for num >= 10 {
		result += "X"
		num -= 10
	}

	for num >= 9 {
		result += "IX"
		num -= 9
	}

	for num >= 5 {
		result += "V"
		num -= 5
	}

	for num >= 4 {
		result += "IV"
		num -= 4
	}

	for num >= 1 {
		result += "I"
		num--
	}

	return result
}

func main() {
	fmt.Println(intToRoman(3))    // III
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(58))   // LVIII
	fmt.Println(intToRoman(1994)) // MCMXCIV
}
