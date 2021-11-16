package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/

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
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
