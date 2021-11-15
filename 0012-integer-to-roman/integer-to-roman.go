package main

import "fmt"

func intToRoman(num int) string {
	result := ""

	for num >= 1000 {
		result += "M"
		num = num - 1000
	}

	for num >= 900 {
		result += "CM"
		num = num - 900
	}

	for num >= 500 {
		result += "D"
		num = num - 500
	}

	for num >= 400 {
		result += "CD"
		num = num - 400
	}

	for num >= 100 {
		result += "C"
		num = num - 100
	}

	for num >= 90 {
		result += "XC"
		num = num - 90
	}

	for num >= 50 {
		result += "L"
		num = num - 50
	}

	for num >= 40 {
		result += "XL"
		num = num - 40
	}

	for num >= 10 {
		result += "X"
		num = num - 10
	}

	for num >= 9 {
		result += "IX"
		num = num - 9
	}

	for num >= 5 {
		result += "V"
		num = num - 5
	}

	for num >= 4 {
		result += "IV"
		num = num - 4
	}

	for num >= 1 {
		result += "I"
		num = num - 1
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
