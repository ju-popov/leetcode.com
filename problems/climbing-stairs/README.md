# 70. Climbing Stairs

**Difficulty**: Easy

**Related Topics**:
- [Math](https://leetcode.com/tag/math/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Memoization](https://leetcode.com/tag/memoization/)

## Problem:

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

**Example 1:**

```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

**Example 2:**

```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

**Constraints:**

- `1 <= n <= 45`

## Solution:

```go
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
```

# Similar Questions:
- [Min Cost Climbing Stairs](https://github.com/ju-popov/leetcode.com/tree/main/problems/min-cost-climbing-stairs/)
- [Fibonacci Number](https://github.com/ju-popov/leetcode.com/tree/main/problems/fibonacci-number/)
- [N-th Tribonacci Number](https://github.com/ju-popov/leetcode.com/tree/main/problems/n-th-tribonacci-number/)
