package main

import "fmt"

/*

58. Length of Last Word

https://leetcode.com/problems/length-of-last-word/

#string

*/

func lengthOfLastWord(s string) int {
	result := 0
	wordFound := false

	for index := len(s) - 1; index >= 0; index-- {
		char := s[index]

		if char == ' ' {
			if !wordFound {
				continue
			}

			return result
		}

		wordFound = true
		result++
	}

	return result
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))                 // 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // 6
}
