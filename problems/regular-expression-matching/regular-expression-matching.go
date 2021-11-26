package main

import "fmt"

/*

10. Regular Expression Matching

https://leetcode.com/problems/regular-expression-matching/

#string #dynamic-programming #recursion

*/

type memoryKey struct {
	s string
	p string
}

func isMatchHelper(s string, p string, memory map[memoryKey]bool) bool {
	if p == "" {
		return s == ""
	}

	key := memoryKey{s: s, p: p}

	if value, ok := memory[key]; ok {
		return value
	}

	symbol := p[0]
	asterisk := len(p) > 1 && p[1] == '*'

	// (dot)*
	if (symbol == '.') && asterisk {
		result := false

		for counter := 0; (!result) && (counter <= len(s)); counter++ {
			result = isMatchHelper(s[counter:], p[2:], memory)
		}

		memory[key] = result

		return result
	}

	// (dot)
	if symbol == '.' {
		if s == "" {
			memory[key] = false

			return false
		}

		result := isMatchHelper(s[1:], p[1:], memory)
		memory[key] = result

		return result
	}

	// (symbol)*
	if asterisk {
		result := isMatchHelper(s, p[2:], memory)

		for counter := 0; (!result) && (counter < len(s)) && (s[counter] == symbol); counter++ {
			result = isMatchHelper(s[counter+1:], p[2:], memory)
		}

		memory[key] = result

		return result
	}

	// (symbol)
	if (s == "") || (s[0] != p[0]) {
		memory[key] = false

		return false
	}

	result := isMatchHelper(s[1:], p[1:], memory)
	memory[key] = result

	return result
}

func isMatch(s string, p string) bool {
	memory := make(map[memoryKey]bool)

	return isMatchHelper(s, p, memory)
}

func main() {
	fmt.Println(isMatch("aa", "a"))                   // false
	fmt.Println(isMatch("aa", "a*"))                  // true
	fmt.Println(isMatch("abc", ".*"))                 // true
	fmt.Println(isMatch("aab", "c*a*b"))              // true
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
	fmt.Println(isMatch("ab", ".*c"))                 // false
}
