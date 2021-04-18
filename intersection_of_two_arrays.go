package main

import "fmt"

/**
https://leetcode.com/problems/intersection-of-two-arrays/

Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.


*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", intersection([]int{1,2,2,1}, []int{2,2})) //[2]
	fmt.Println("Answer is:", intersection([]int{4,9,5}, []int{9,4,9,8,4})) //[9,4]

	fmt.Println("Answer is:", intersectionRetainAll([]int{1,2,2,1}, []int{2,2})) //[2,2]
	fmt.Println("Answer is:", intersectionRetainAll([]int{4,9,9,5}, []int{9,4,9,8,4})) //[9,4,9]
}

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	checkMap := make(map[int]bool)

	for _, num := range nums1 {
		checkMap[num] = true
	}
	for _, num := range nums2 {
		if checkMap[num] {
			result = append(result, num)
			checkMap[num] = false
		}
	}
	return result
}

/**
A slight variation: https://leetcode.com/problems/intersection-of-two-arrays-ii/

Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.
 */

func intersectionRetainAll(nums1 []int, nums2 []int) []int {
	var result []int
	checkMap := make(map[int]int)

	for _, num := range nums1 {
		checkMap[num] ++
	}
	for _, num := range nums2 {
		if checkMap[num] > 0 {
			result = append(result, num)
			checkMap[num] --
		}
	}
	return result
}