package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

Given an integer array nums, you need to find one continuous subarray that if you only
sort this subarray in ascending order, then the whole array will be sorted in ascending order.

Return the shortest such subarray and output its length.

Example 1:
Input: nums = [2,6,4,8,10,9,15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.

Example 2:
Input: nums = [1,2,3,4]
Output: 0

Example 3:
Input: nums = [1]
Output: 0
 */

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findUnsortedSubarray([]int{2,6,4,8,10,9,15})) //5
	fmt.Println("Answer is:", findUnsortedSubarray([]int{1,2,3,4})) //0
	fmt.Println("Answer is:", findUnsortedSubarray([]int{1})) //0
}

/**
 easy solution, O(nlogn)
 */
func findUnsortedSubarray(nums []int) int {
	sortedArray := make([]int, len(nums))
	copy(sortedArray, nums)
	sort.Ints(sortedArray)

	min, max := -1, -2

	//compare the sorted with unsorted. the element at same index "i" should be the same if it's sorted
	for i, num := range nums {
		if num != sortedArray[i] {
			if min < 0 { min = i }
			max = i
		}
	}
	return max-min+1
}

/**
 O(n) solution
 */
func findUnsortedSubarray_Optimal(nums []int) int {


	return 0
}