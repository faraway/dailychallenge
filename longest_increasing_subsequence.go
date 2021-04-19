package main

import (
	"fmt"
	"math"
)

/**
Longest Increasing Subsequence (LIS)
https://leetcode.com/problems/longest-increasing-subsequence/

Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements
without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:
1 <= nums.length <= 2500
-104 <= nums[i] <= 104

Follow up:
Could you come up with the O(n2) solution?
Could you improve it to O(n log(n)) time complexity?
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", lengthOfLIS_bruteforce([]int{10,9,2,5,3,7,101,18})) //4
	fmt.Println("Answer is:", lengthOfLIS_bruteforce([]int{0,1,0,3,2,3})) //4
	fmt.Println("Answer is:", lengthOfLIS_bruteforce([]int{7,7,7,7,7,7,7})) //1

	fmt.Println("Answer is:", lengthOfLIS([]int{10,9,2,5,3,7,101,18})) //4
	fmt.Println("Answer is:", lengthOfLIS([]int{0,1,0,3,2,3})) //4
	fmt.Println("Answer is:", lengthOfLIS([]int{7,7,7,7,7,7,7})) //1
}

/**
 Brute force algorithm, time complexity can be up to 2^n, where n is the length of input array
 */
type Option struct {
	Length int
	Max int
}

func lengthOfLIS_bruteforce(nums []int) int {
	options := []Option {{0, math.MinInt32}}
	for _, num := range nums {
		var newOptions []Option
		for _, option := range options {
			if num > option.Max {
				newOptions = append(newOptions, Option{option.Length+1, num})
			}
		}
		options = append(options, newOptions...)
	}

	result := 0
	for _, option := range options {
		if option.Length > result {
			result = option.Length
		}
	}
	return result
}

/**
Improved solution, but kinda a improvement over the brute force solution above
*/
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	/**
		result is the array for DP.
		result[i] means the smallest possible value for subsequence of length i
		result[0] set to min value to get an easy start
	 */
	result := []int{math.MinInt32}
	for _, num := range nums {
		currLen := len(result)
		//if current number > the last number of current sequence, then we can extend the sequence length by 1
		if num > result[currLen-1] {
			result = append(result, num)
		}
		//update the result array. If current number is larger than last value of sequence length=i-1 and smaller than last value of sequence length=i
		//then last value of sequence length=i can be updated to current number.
		for i := currLen-2; i>=0; i-- {
			if num > result[i] && num < result[i+1] {
				result[i+1] = num
			}
		}
	}

	return len(result) - 1
}

//TODO - read, another DP but slightly different solution:
//https://github.com/faraway/fucking-algorithm/blob/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E8%AE%BE%E8%AE%A1%EF%BC%9A%E6%9C%80%E9%95%BF%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97.md
