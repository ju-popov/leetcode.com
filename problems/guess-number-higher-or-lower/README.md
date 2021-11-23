# 374. Guess Number Higher or Lower

**Difficulty**: Easy

## Related Topics:

- [Binary Search](https://leetcode.com/tag/binary-search/)
- [Interactive](https://leetcode.com/tag/interactive/)

## Problem:

We are playing the Guess Game. The game is as follows:

I pick a number from `1` to `n`. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API `int guess(int num)`, which returns 3 possible results:

- `-1`: The number I picked is lower than your guess (i.e. `pick < num`).
- `1`: The number I picked is higher than your guess (i.e. `pick > num`).
- `0`: The number I picked is equal to your guess (i.e. `pick == num`).

Return *the number that I picked*.

**Example 1:**

```
Input: n = 10, pick = 6
Output: 6
```

**Example 2:**

```
Input: n = 1, pick = 1
Output: 1
```

**Example 3:**

```
Input: n = 2, pick = 1
Output: 1
```

**Example 4:**

```
Input: n = 2, pick = 2
Output: 2
```

**Constraints:**

- `1 <= n <= 231 - 1`
- `1 <= pick <= n`

## Solution:

```go
func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := left + (right-left)/2

		switch guess(mid) {
		case 1:
			left = mid + 1
		case -1:
			right = mid - 1
		default:
			return mid
		}
	}

	return -1
}
```

## Similar Questions:

- [First Bad Version](https://github.com/ju-popov/leetcode.com/tree/main/problems/first-bad-version/)
- [Guess Number Higher or Lower II](https://github.com/ju-popov/leetcode.com/tree/main/problems/guess-number-higher-or-lower-ii/)
- [Find K Closest Elements](https://github.com/ju-popov/leetcode.com/tree/main/problems/find-k-closest-elements/)

