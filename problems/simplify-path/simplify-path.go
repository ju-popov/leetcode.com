package main

import (
	"fmt"
	"strings"
)

/*

71. Simplify Path

https://leetcode.com/problems/simplify-path/

#string #stack

*/

func simplifyPath(path string) string {
	stack := make([]string, 0, 1000)
	parts := strings.Split(path, "/")

	for _, part := range parts {
		// Nothing
		if len(part) == 0 {
			continue
		}

		// Current directory
		if part == "." {
			continue
		}

		// Parent directory
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

			continue
		}

		// Directory or file
		stack = append(stack, part)
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))          // /home
	fmt.Println(simplifyPath("/../"))            // /
	fmt.Println(simplifyPath("/home//foo/"))     // /home/foo
	fmt.Println(simplifyPath("/a/./b/../../c/")) // /c
}
