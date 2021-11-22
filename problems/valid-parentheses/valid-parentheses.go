package main

/*

20. Valid Parentheses

https://leetcode.com/problems/valid-parentheses/

#string #stack

*/

import "fmt"

var parentheses = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if value, ok := parentheses[char]; ok {
			// opening
			stack = append(stack, value)
		} else {
			// closing
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != char {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
