package main

import "fmt"

/**
https://leetcode.com/problems/subarray-sum-equals-k/

Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2

Example 2:
Input: nums = [1,2,3], k = 3
Output: 2


Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", subarraySum([]int{1,1,1}, 2)) //2
	fmt.Println("Answer is:", subarraySum([]int{1,2,3}, 3)) //2

}

func subarraySum(nums []int, k int) int {
	count := 0

	sum := 0
	//key: prefix sum; value: count of such prefix sum (because there could be negative nums)
	//"prefix sum" is sum[0,i] where i is the current index
	sumMap := make(map[int]int)

	for _, num := range nums {
		sum += num
		if sum == k { //array sum from beginning
			count++
		}
		occurrence, exists := sumMap[sum-k] //subarray starts somewhere in the middle
		if exists {
			count += occurrence
		}
		sumMap[sum]++
	}
	return count
}
