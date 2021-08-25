package main

import "fmt"

/**
https://leetcode.com/problems/continuous-subarray-sum/
Given an integer array nums and an integer k,
return true if nums has a continuous subarray of size at least two whose elements sum up to a multiple of k, or false otherwise.

An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.

Example 1:
Input: nums = [23,2,4,6,7], k = 6
Output: true
Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

Example 2:
Input: nums = [23,2,6,4,7], k = 6
Output: true
Explanation: [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.

Example 3:
Input: nums = [23,2,6,4,7], k = 13
Output: false

Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 109
0 <= sum(nums[i]) <= 231 - 1
1 <= k <= 231 - 1

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", checkSubarraySum([]int{23,2,4,6,7}, 6)) //True
	fmt.Println("Answer is:", checkSubarraySum([]int{23,2,6,4,7}, 6)) //True
	fmt.Println("Answer is:", checkSubarraySum([]int{23,2,6,4,7}, 13)) //False
	fmt.Println("Answer is:", checkSubarraySum([]int{1,0}, 2)) //False
	fmt.Println("Answer is:", checkSubarraySum([]int{1,0,0}, 2)) //True. 0 is always a multiplier of K as long as subarray size >=2
}

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	sum := 0 // prefix sum
	sumMap := make(map[int]int) //key:sum (mod by k); value: index of this prefix sum

	//brute force - will timeout
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum %= k

		//the continuous array starts from beginning, i>0 because we need at least 2 elements as defined by the problem
		if sum == 0 && i > 0 {
			return true
		}
		//the continuous array starts from somewhere in the middle. so then current_sum-k should equal to that prefix sum.
		//since we took sum=sum%k, then current_sum-k should simply just be same current_sum, hence the sumMap[sum] check
		idx, exists := sumMap[sum]
		if exists && (i - idx > 1) {
			return true
		}
		if !exists {
			sumMap[sum] = i
		}
	}

	return false
}

func checkSubarraySum_brute_force(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	for i, num := range nums {
		nums[i] = num % k
	}

	//brute force - will timeout
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i+1; j < len(nums); j++ {
			sum += nums[j]
			if sum % k == 0 {
				return true
			}
		}
	}

	return false
}

