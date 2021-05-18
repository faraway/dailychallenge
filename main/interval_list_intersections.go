package main

import "fmt"

/**
https://leetcode.com/problems/interval-list-intersections/

You are given two lists of closed intervals, firstList and secondList,
where firstList[i] = [starti, endi] and secondList[j] = [startj, endj].
Each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

A closed interval [a, b] (with a < b) denotes the set of real numbers x with a <= x <= b.

The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval.
For example, the intersection of [1, 3] and [2, 4] is [2, 3].


Example 1:
Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

Example 2:
Input: firstList = [[1,3],[5,9]], secondList = []
Output: []

Example 3:
Input: firstList = [], secondList = [[4,8],[10,12]]
Output: []

Example 4:
Input: firstList = [[1,7]], secondList = [[3,10]]
Output: [[3,7]]


Constraints:
0 <= firstList.length, secondList.length <= 1000
firstList.length + secondList.length >= 1
0 <= starti < endi <= 109
endi < starti+1
0 <= startj < endj <= 109
endj < startj+1

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", intervalIntersection([][]int{{0,2},{5,10},{13,23},{24,25}},
								[][]int{{1,5},{8,12},{15,24},{25,26}})) //[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
	fmt.Println("Answer is:", intervalIntersection([][]int{{1,3},{5,9}}, [][]int{})) //[]
	fmt.Println("Answer is:", intervalIntersection([][]int{}, [][]int{{4,8},{10,12}})) //[]
	fmt.Println("Answer is:", intervalIntersection([][]int{{1,7}}, [][]int{{3,10}})) //[[3,7]]
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int

	first, second := 0,0
	for first < len(firstList) && second < len(secondList) {
		interval1 := firstList[first]
		interval2 := secondList[second]

		//compare 'end' time of the two intervals
		if interval1[1] < interval2[1] {
			if interval2[0] <= interval1[0] {
				/**
				  interval1      |----------|
				  interval2    |----------------|
				*/
				result = append(result, []int{interval1[0], interval1[1]})
			} else if interval2[0] >= interval1[0] && interval2[0] <= interval1[1] {
				/**
				  interval1      |----------|
				  interval2            |--------|
				*/
				result = append(result, []int{interval2[0], interval1[1]})
			}
			/**
			      interval1      |----------|
			      interval2                   |-|
			      (no overlap, ignore this situation
			*/
			first++
		} else { //same as above, but interval1 and interval2 in opposite situations
			if interval1[0] <= interval2[0] {
				result = append(result, []int{interval2[0], interval2[1]})
			} else if interval1[0] >= interval2[0] && interval1[0] <= interval2[1] {
				result = append(result, []int{interval1[0], interval2[1]})
			}
			second++
		}
	}
	return result
}
