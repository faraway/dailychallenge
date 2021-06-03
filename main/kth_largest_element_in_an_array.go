package main

import "fmt"

/**
https://leetcode.com/problems/kth-largest-element-in-an-array/solution/

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.


Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

Constraints:

1 <= k <= nums.length <= 104
-104 <= nums[i] <= 104

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findKthLargest([]int{3,2,1,5,6,4}, 2)) //5
	fmt.Println("Answer is:", findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4)) //4
}

/**
 We can use a minHeap holding k largest elements of the array. It will be O(Nlogk)
 here we use "quickselect" algorithm which is said to has an average performance of O(n)
 */
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k int, start int, end int) int {

	pivotIndex := partition(nums, start, end)
	//fmt.Println("pivot index is:", pivotIndex, "; array is:", nums)

	//partition is based on small elements on left and large elements on right
	//kth largst ideally should = n - pivotIndex
	if len(nums) - pivotIndex == k {
		return nums[pivotIndex]
	} else if len(nums) - pivotIndex > k {
		return quickSelect(nums, k, pivotIndex+1, end)
	} else {
		return quickSelect(nums, k, start, pivotIndex-1)
	}
}

/**
	Partition the array[start:end] (both inclusive) using array[end] as pivot
	return the index of the pivot after partition
 */
func partition(nums []int, start int, end int) int {
	pivot := nums[end]
	left, right := start, end-1
	for left <= right {
		if nums[left] < pivot {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			nums[right], nums[right+1] = nums[right+1], nums[right] //swap with pivot
			right--
		}
	}
	return left
}