package main

/*

12. Integer to Roman

https://leetcode.com/problems/integer-to-roman/

#hash-table #math #string

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
