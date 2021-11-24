package main

import "fmt"

/*

38. Count and Say

https://leetcode.com/problems/count-and-say/

#string

*/

func countAndSayHelper(input string) string {
	if input == "" {
		return "1"
	}

	count := 0
	result := ""

	var lastChar byte

	for index := 0; index < len(input); index++ {
		char := input[index]

		if index == 0 {
			lastChar = char
			count = 1

			continue
		}

		if char == lastChar {
			count++

			continue
		}

		result += fmt.Sprintf("%d%s", count, string(lastChar))

		lastChar = char
		count = 1
	}

	result += fmt.Sprintf("%d%s", count, string(lastChar))

	return result
}

func countAndSay(n int) string {
	result := ""

	for i := 0; i < n; i++ {
		result = countAndSayHelper(result)
	}

	return result
}

func main() {
	fmt.Println(countAndSay(1))  // 1
	fmt.Println(countAndSay(4))  // 1211
	fmt.Println(countAndSay(10)) // 13211311123113112211
}
