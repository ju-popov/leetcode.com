package main

/*

20. Valid Parentheses

https://leetcode.com/problems/valid-parentheses/

Approach 1: Stacks

Complexity analysis

Time complexity : O(n) because we simply traverse the given string one character
at a time and push and pop operations on a stack take O(1) time.

Space complexity : O(n) as we push all opening brackets onto the stack and in
the worst case, we will end up pushing all the brackets onto the stack.
e.g. ((((((((((.

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
