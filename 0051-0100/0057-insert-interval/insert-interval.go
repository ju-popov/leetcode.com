package main

import (
	"fmt"
	"sort"
)

/*

57. Insert Interval

https://leetcode.com/problems/insert-interval/

*/

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
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results := make([][]int, 0)
	results = append(results, intervals[0])

	for index := 1; index < len(intervals); index++ {
		interval := intervals[index]

		if isOverlapping(interval, results[len(results)-1]) {
			results[len(results)-1] = mergeIntervals(interval, results[len(results)-1])

			continue
		}

		results = append(results, interval)
	}

	return results
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})) // [[1 5] [6 9]]
	fmt.Println(insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}),
	) // [[1 2] [3 10] [12 16]]
	fmt.Println(insert([][]int{}, []int{5, 7}))                 // [[5 7]]
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}))           // [[1 5]]
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))           // [[1 7]]
	fmt.Println(insert([][]int{{1, 5}}, []int{6, 8}))           // [[1 5] [6 8]]
	fmt.Println(insert([][]int{{3, 5}, {12, 15}}, []int{6, 6})) // [[3 5] [6 6] [12 15]]
}
