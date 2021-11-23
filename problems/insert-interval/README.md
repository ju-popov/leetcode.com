# 57. Insert Interval

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)

## Problem:

You are given an array of non-overlapping intervals `intervals` where `intervals[i] = [starti, endi]` represent the start and the end of the `ith` interval and `intervals` is sorted in ascending order by `starti`. You are also given an interval `newInterval = [start, end]` that represents the start and end of another interval.

Insert `newInterval` into `intervals` such that `intervals` is still sorted in ascending order by `starti` and `intervals` still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return `intervals`*after the insertion*.

**Example 1:**

```
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
```

**Example 2:**

```
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
```

**Example 3:**

```
Input: intervals = [], newInterval = [5,7]
Output: [[5,7]]
```

**Example 4:**

```
Input: intervals = [[1,5]], newInterval = [2,3]
Output: [[1,5]]
```

**Example 5:**

```
Input: intervals = [[1,5]], newInterval = [2,7]
Output: [[1,7]]
```

**Constraints:**

- `0 <= intervals.length <= 104`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 105`
- `intervals` is sorted by `starti` in **ascending** order.
- `newInterval.length == 2`
- `0 <= start <= end <= 105`

## Solution:

```go
func minInt(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}

func maxInt(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func isOverlapping(interval1 []int, interval2 []int) bool {
	if interval1[1] < interval2[0] {
		return false
	}

	if interval1[0] > interval2[1] {
		return false
	}

	return true
}

func mergeIntervals(interval1 []int, interval2 []int) []int {
	return []int{
		minInt(interval1[0], interval2[0]),
		maxInt(interval1[1], interval2[1]),
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	results := make([][]int, 0)

	index := 0
	for index < len(intervals) && intervals[index][0] < newInterval[0] {
		results = append(results, intervals[index])
		index++
	}

	if (len(results) > 0) && isOverlapping(newInterval, results[len(results)-1]) {
		results[len(results)-1] = mergeIntervals(newInterval, results[len(results)-1])
	} else {
		results = append(results, newInterval)
	}

	for index < len(intervals) {
		if isOverlapping(intervals[index], results[len(results)-1]) {
			results[len(results)-1] = mergeIntervals(intervals[index], results[len(results)-1])
		} else {
			results = append(results, intervals[index])
		}
		index++
	}

	return results
}
```

## Similar Questions:

- [Merge Intervals](https://github.com/ju-popov/leetcode.com/tree/main/problems/merge-intervals/)
- [Range Module](https://github.com/ju-popov/leetcode.com/tree/main/problems/range-module/)
