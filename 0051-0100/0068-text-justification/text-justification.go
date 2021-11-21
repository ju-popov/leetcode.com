package main

import (
	"fmt"
	"strings"
)

/*

68. Text Justification

https://leetcode.com/problems/text-justification/

*/

func nextWords(words []string, maxWidth int, index int) ([]string, int) {
	result := make([]string, 0)

	width := 0
	first := true

	for {
		if index >= len(words) {
			return result, index
		}

		word := words[index]

		width += len(word)
		if !first {
			width++
		} else {
			first = false
		}

		if width > maxWidth {
			return result, index
		}

		result = append(result, word)

		index++
	}
}

func generateRow(words []string, maxWidth int, lastRow bool) string {
	if lastRow {
		result := strings.Join(words, " ")

		return result + strings.Repeat(" ", maxWidth-len(result))
	}

	totalWordsWidth := 0
	for _, word := range words {
		totalWordsWidth += len(word)
	}

	result := ""
	for index, word := range words {
		result += word
		totalWordsWidth -= len(word)
		extraspace := maxWidth - len(result) - totalWordsWidth
		wordsLeft := len(words) - index - 1

		if wordsLeft == 0 {
			result += strings.Repeat(" ", maxWidth-len(result))
		} else {
			result += strings.Repeat(" ", extraspace/wordsLeft)
			if extraspace%wordsLeft != 0 {
				result += " "
			}
		}
	}

	return result
}

func fullJustify(words []string, maxWidth int) []string {
	results := make([]string, 0)

	index := 0

	for {
		var row []string
		row, index = nextWords(words, maxWidth, index)

		results = append(results, generateRow(row, maxWidth, index >= len(words)))

		if index >= len(words) {
			break
		}
	}

	return results
}

func printResult(result []string) {
	for _, row := range result {
		fmt.Println("[" + row + "]")
	}

	fmt.Println()
}

func main() {
	words1 := []string{
		"This",
		"is",
		"an",
		"example",
		"of",
		"text",
		"justification.",
	}

	// [This    is    an]
	// [example  of text]
	// [justification.  ]

	printResult(fullJustify(words1, 16))

	words2 := []string{
		"What",
		"must",
		"be",
		"acknowledgment",
		"shall",
		"be",
	}

	// [What   must   be]
	// [acknowledgment  ]
	// [shall be        ]

	printResult(fullJustify(words2, 16))

	words3 := []string{
		"Science",
		"is",
		"what",
		"we",
		"understand",
		"well",
		"enough",
		"to",
		"explain",
		"to",
		"a",
		"computer.",
		"Art",
		"is",
		"everything",
		"else",
		"we",
		"do",
	}

	// [Science  is  what we]
	// [understand      well]
	// [enough to explain to]
	// [a  computer.  Art is]
	// [everything  else  we]
	// [do                  ]

	printResult(fullJustify(words3, 20))
}
