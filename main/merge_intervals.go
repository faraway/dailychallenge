package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/merge-intervals/

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", merge([][]int{{1,9}})) //[[1,9]]
	fmt.Println("Answer is:", merge([][]int{{1,3},{2,6},{8,10},{15,18}})) //[[1,6],[8,10],[15,18]]
	fmt.Println("Answer is:", merge([][]int{{1,4},{4,5}})) //[[1,5]]

}

func merge(intervals [][]int) [][]int {
	var result [][]int
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0] //sort by each interval's start time
	})

	result = append(result, intervals[0])
	for i:=1; i<len(intervals); i++ {
		previousInterval := result[len(result)-1]
		if previousInterval[1] >= intervals[i][0] {
			if previousInterval[1] < intervals[i][1] {
				previousInterval[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}