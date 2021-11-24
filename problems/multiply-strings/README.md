# 43. Multiply Strings

**Difficulty**: Medium

## Related Topics:

- [Math](https://leetcode.com/tag/math/)
- [String](https://leetcode.com/tag/string/)
- [Simulation](https://leetcode.com/tag/simulation/)

## Problem:

Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string.

**Note:** You must not use any built-in BigInteger library or convert the inputs to integer directly.

**Example 1:**

```
Input: num1 = "2", num2 = "3"
Output: "6"
```

**Example 2:**

```
Input: num1 = "123", num2 = "456"
Output: "56088"
```

**Constraints:**

- `1 <= num1.length, num2.length <= 200`
- `num1` and `num2` consist of digits only.
- Both `num1` and `num2` do not contain any leading zero, except the number `0` itself.

## Solution:

```go
func addToResult(result []byte, value byte, offset int) {
	for index := len(result) - 1 - offset; index >= 0; index-- {
		if value == 0 {
			return
		}

		total := result[index] - '0' + value

		result[index] = total%10 + '0'

		value = total / 10
	}
}

func multiply(num1 string, num2 string) string {
	// Create result buffer with zeroes
	result := make([]byte, len(num1)+len(num2))
	for index := 0; index < len(result); index++ {
		result[index] = '0'
	}

	// Multiply numbers
	for index2 := len(num2) - 1; index2 >= 0; index2-- {
		for index1 := len(num1) - 1; index1 >= 0; index1-- {
			n1 := num1[index1] - '0'
			n2 := num2[index2] - '0'
			offset := (len(num2) - 1 - index2) + (len(num1) - 1 - index1)

			addToResult(result, n1*n2, offset)
		}
	}

	// Strip leading zeroes
	var index int
	for index = 0; index < len(result)-1; index++ {
		if result[index] != '0' {
			break
		}
	}

	return string(result[index:])
}
```

## Similar Questions:

- [Add Two Numbers](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-two-numbers/)
- [Plus One](https://github.com/ju-popov/leetcode.com/tree/main/problems/plus-one/)
- [Add Binary](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-binary/)
- [Add Strings](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-strings/)
