package main

import "fmt"

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
}

/**
 Brute force algorithm, time complexity can be up to 2^n, where n is the length of input array
 */
func lengthOfLIS_bruteforce(nums []int) int {
	options := []Option {{0, -9999}}
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

type Option struct {
	Length int
	Max int
}