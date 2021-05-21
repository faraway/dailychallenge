package main

import (
	"fmt"
)

/**
https://leetcode.com/problems/search-in-rotated-sorted-array/

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the rotation and an integer target, return the index of target if it is in nums,
or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is guaranteed to be rotated at some pivot.
-104 <= target <= 104

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", search([]int{4,5,6,7,0,1,2}, 0)) //4
	fmt.Println("Answer is:", search([]int{4,5,6,7,0,1,2}, 3)) //-1
	fmt.Println("Answer is:", search([]int{1}, 0)) //-1
	fmt.Println("Answer is:", search([]int{3,1}, 1)) //1
}


func search(nums []int, target int) int {
	return searchRecursive(nums, target, 0, len(nums)-1)
}

func searchRecursive(nums []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end - start)/2
	if nums[mid] == target {
		return mid
	}

	if nums[start] <= nums[end] { //fully sorted array, this is normal binary search. We probably don't need this, sort of optimization
		if target < nums[mid] {
			return searchRecursive(nums, target, start, mid-1)
		} else {
			return searchRecursive(nums, target, mid+1, end)
		}
	} else if nums[start] <= nums[mid] { // left half is sorted, pivot point is in right half
		if nums[start] <= target && target <= nums[mid] {
			return searchRecursive(nums, target, start, mid-1)
		} else {
			return searchRecursive(nums, target, mid+1, end)
		}
	} else { //right half is sorted
		if nums[mid] <= target && target <= nums[end] {
			return searchRecursive(nums, target, mid+1, end)
		} else {
			return searchRecursive(nums, target, start, mid-1)
		}
	}
}