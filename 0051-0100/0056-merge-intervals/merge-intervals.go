package main

import (
	"fmt"
	"sort"
)

/*

56. Merge Intervals

https://leetcode.com/problems/merge-intervals/

*/

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

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1 5]]
}
