package main

import (
	"container/heap"
	"fmt"
)

/**
https://leetcode.com/problems/k-closest-points-to-origin/

Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k,
return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)^2 + (y1 - y2)^2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).



Example 1:
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].

Example 2:
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", kClosest([][]int{{1,3},{-2,2}}, 1)) //[[-2,2]]
	fmt.Println("Answer is:", kClosest([][]int{{3,3},{5,-1},{-2,4}}, 2)) //[[3,3],[-2,4]] or [[-2 4] [3 3]]
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxPointHeap{}
	heap.Init(h)

	for _, point := range points {
		if h.Len() < k {
			heap.Push(h, point)
		} else if distanceLargerThan((*h)[0],point) {
			heap.Pop(h)
			heap.Push(h, point)
		}
	}

	return *h
}

func distanceLargerThan(a []int, b[]int) bool {
	return (a[0]*a[0] + a[1]*a[1]) > (b[0]*b[0] + b[1]*b[1])
}

type MaxPointHeap [][]int

func (m MaxPointHeap) Len() int {
	return len(m)
}

func (m MaxPointHeap) Less(i, j int) bool {
	return distanceLargerThan(m[i], m[j]) //max heap
}

func (m MaxPointHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxPointHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *MaxPointHeap) Pop() interface{} {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}


