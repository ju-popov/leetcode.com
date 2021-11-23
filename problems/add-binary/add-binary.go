package main

import (
	"fmt"
)

/*

67. Add Binary

https://leetcode.com/problems/add-binary/

#math #string #bit-manipulation #simulation

*/

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func addBinary(a string, b string) string {
	result := make([]byte, max(len(a), len(b)))
	indexA := len(a) - 1
	indexB := len(b) - 1
	indexResult := len(result) - 1
	carry := byte(0)

	for (indexA >= 0) || (indexB >= 0) {
		sum := carry

		if indexA >= 0 {
			sum += a[indexA] - '0'
		}

		if indexB >= 0 {
			sum += b[indexB] - '0'
		}

		carry = sum / 2
		result[indexResult] = '0' + sum%2

		indexA--
		indexB--
		indexResult--
	}

	if carry > 0 {
		return "1" + string(result)
	}

	return string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("1010", "1011")) // 10101
	fmt.Println(addBinary("0", "0"))       // 10101
}
