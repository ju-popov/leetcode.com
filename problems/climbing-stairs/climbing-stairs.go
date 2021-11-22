package main

import "fmt"

/*

70. Climbing Stairs

https://leetcode.com/problems/climbing-stairs/

#math #dynamic-programming #memoization

*/

func climbStairsHelper(n int, memory map[int]int) int {
	if (n == 1) || (n == 2) {
		return n
	}

	if value, ok := memory[n]; ok {
		return value
	}

	oneStepNext := climbStairsHelper(n-1, memory)
	twoStepsNext := climbStairsHelper(n-2, memory)

	total := oneStepNext + twoStepsNext

	memory[n] = total

	return total
}

func climbStairs(n int) int {
	memory := make(map[int]int)

	return climbStairsHelper(n, memory)
}

func main() {
	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(3)) // 3
}
