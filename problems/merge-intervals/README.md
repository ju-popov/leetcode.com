# 56. Merge Intervals

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Sorting](https://leetcode.com/tag/sorting/)

## Problem:

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return *an array of the non-overlapping intervals that cover all the intervals in the input*.

**Example 1:**

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

**Example 2:**

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

**Constraints:**

- `1 <= intervals.length <= 104`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 104`

## Solution:

```go
func minInt(v1 int, v2 int) int {
	if v1 <= v2 {
		return v1
	}

	return v2
}

func maxInt(v1 int, v2 int) int {
	if v1 >= v2 {
		return v1
	}

	return v2
}

func merge(intervals [][]int) [][]int {
	results := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastInterval := intervals[0]

	for index := 1; index < len(intervals); index++ {
		interval := intervals[index]

		if interval[0] <= lastInterval[1] {
			lastInterval[0] = minInt(lastInterval[0], interval[0])
			lastInterval[1] = maxInt(lastInterval[1], interval[1])

			continue
		}

		results = append(results, lastInterval)
		lastInterval = interval
	}

	results = append(results, lastInterval)

	return results
}
```

## Similar Questions:

- [Insert Interval](https://github.com/ju-popov/leetcode.com/tree/main/problems/insert-interval/)
- [Meeting Rooms](https://github.com/ju-popov/leetcode.com/tree/main/problems/meeting-rooms/)
- [Meeting Rooms II](https://github.com/ju-popov/leetcode.com/tree/main/problems/meeting-rooms-ii/)
- [Teemo Attacking](https://github.com/ju-popov/leetcode.com/tree/main/problems/teemo-attacking/)
- [Add Bold Tag in String](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-bold-tag-in-string/)
- [Range Module](https://github.com/ju-popov/leetcode.com/tree/main/problems/range-module/)
- [Employee Free Time](https://github.com/ju-popov/leetcode.com/tree/main/problems/employee-free-time/)
- [Partition Labels](https://github.com/ju-popov/leetcode.com/tree/main/problems/partition-labels/)
- [Interval List Intersections](https://github.com/ju-popov/leetcode.com/tree/main/problems/interval-list-intersections/)
