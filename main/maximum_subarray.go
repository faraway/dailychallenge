package main

import "fmt"

/**
https://leetcode.com/problems/maximum-subarray/ (EASY)

Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.


Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

Constraints:

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})) //6
	fmt.Println("Answer is:", maxSubArray([]int{1})) //1
	fmt.Println("Answer is:", maxSubArray([]int{-3})) //-3
	fmt.Println("Answer is:", maxSubArray([]int{5,4,-1,7,8})) //23
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			//discard previous number(s) as this accumulative sum will be contributing negatively to the subarray
			//and start from current number as the new beginning of the potential max subarray
			currentSum = 0
		}
		currentSum += nums[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}