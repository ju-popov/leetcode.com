package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	results := []string{""}

	for index := 0; index < 2*n; index++ {
		newResults := make([]string, 0)

		for _, result := range results {
			opening := strings.Count(result, "(")
			closing := len(result) - opening

			if opening < n {
				newResults = append(newResults, result+"(")
			}

			if opening > closing {
				newResults = append(newResults, result+")")
			}
		}

		results = newResults
	}

	return results
}

func main() {
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}
