package main

import "fmt"

var digitsMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	results := make([]string, 0)
	if len(digits) == 0 {
		return results
	}

	// Add one empty result
	results = append(results, []string{""}...)

	for index := len(digits) - 1; index >= 0; index-- {
		digit := digits[index]
		newResults := make([]string, 0)

		for _, char := range digitsMap[digit] {
			for _, result := range results {
				newResults = append(newResults, char+result)
			}
		}

		results = newResults
	}

	return results
}

func main() {
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations(""))   // []
	fmt.Println(letterCombinations("2"))  // ["a","b","c"]
}
