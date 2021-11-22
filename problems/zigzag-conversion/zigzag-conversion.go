package main

/*

6. Zigzag Conversion

https://leetcode.com/problems/zigzag-conversion/

#string

*/

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	rows := make([]string, numRows)

	goingDown := false
	row := 0

	for _, char := range s {
		rows[row] += string(char)

		if (row == 0) || (row == numRows-1) {
			goingDown = !goingDown
		}

		if goingDown {
			row++
		} else {
			row--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) // PINALSIGYAHRPI
	fmt.Println(convert("A", 1))              // A
}
