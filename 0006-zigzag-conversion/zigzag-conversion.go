package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	rows := make([]string, numRows)

	direction := false
	row := 0

	for _, char := range s {
		rows[row] += string(char)
		if (row == 0) || (row == numRows-1) {
			direction = !direction
		}

		if direction {
			row++
		} else {
			row--
		}
	}

	return strings.Join(rows, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) // PINALSIGYAHRPI
	fmt.Println(convert("A", 1))              // A
}
