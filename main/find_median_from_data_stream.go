package main

import (
	"container/heap"
	. "dailychallenge/utils"
	"fmt"
)

/**
https://leetcode.com/problems/find-median-from-data-stream/

The median is the middle value in an ordered integer list.
If the size of the list is even, there is no middle value and the median is the mean of the two middle values.

For example, for arr = [2,3,4], the median is 3.
For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.


Example 1:

Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0


Constraints:

-10^5 <= num <= 10^5
There will be at least one element in the data structure before calling findMedian.
At most 5 * 10^4 calls will be made to addNum and findMedian.

*/

func main() {
	fmt.Println("initializing test data...")

	obj := newMedianFinder()
	obj.AddNum(5)
	fmt.Println("Answer is:", obj.FindMedian()) //5
	obj.AddNum(4)
	fmt.Println("Answer is:", obj.FindMedian()) //4.5
	obj.AddNum(3)
	obj.AddNum(2)
	obj.AddNum(1)
	fmt.Println("Answer is:", obj.FindMedian()) //3
}

type MedianFinder struct {
	minHeap *IntHeap
	maxHeap *IntMaxHeap
}


/** initialize your data structure here. */
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func newMedianFinder() MedianFinder {
	instance := MedianFinder{minHeap: &IntHeap{}, maxHeap: &IntMaxHeap{}}
	heap.Init(instance.minHeap)
	heap.Init(instance.maxHeap)
	return instance
}


func (this *MedianFinder) AddNum(num int)  {
	maxHeapSize := len(*this.maxHeap)
	minHeapSize := len(*this.minHeap)

	if maxHeapSize == 0 {
		heap.Push(this.maxHeap, num)
		return
	}

	firstHalfMax := (*this.maxHeap)[0] //peek the number but not "pop" it

	if num <= firstHalfMax {
		heap.Push(this.maxHeap, num)
		maxHeapSize++
		if maxHeapSize > minHeapSize+1 {
			v := heap.Pop(this.maxHeap)
			heap.Push(this.minHeap, v)
		}
	} else {
		heap.Push(this.minHeap, num)
		minHeapSize++
		if minHeapSize > maxHeapSize {
			v := heap.Pop(this.minHeap)
			heap.Push(this.maxHeap, v)
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	maxHeapSize := len(*this.maxHeap)
	minHeapSize := len(*this.minHeap)

	median := 0.0
	if maxHeapSize == minHeapSize + 1 { //odd number of elements
		median = float64((*this.maxHeap)[0])
	} else if maxHeapSize == minHeapSize { //even number of elements
		a := float64((*this.maxHeap)[0])
		b := float64((*this.minHeap)[0])
		median = a + (b-a)/2
	} else {
		fmt.Println("you should not be here.")
	}
	return median
}




