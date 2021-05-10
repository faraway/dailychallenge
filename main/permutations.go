package main

import "fmt"

/**
https://leetcode.com/problems/permutations/

Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.



Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]


Constraints:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", permute([]int{1,2,3})) //[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	fmt.Println("Answer is:", permute([]int{0,1})) //[[0,1],[1,0]]
	fmt.Println("Answer is:", permute([]int{1})) //[[1]]
}

func permute(nums []int) [][]int {
     results := make([][]int,0,10)
     backtrack(nums, make([]int,0,len(nums)), &results)
     return results
}

func backtrack(nums []int, track []int, results *[][]int) {
	//add result to collection
	if len(track) == len(nums) {
		result := make([]int, len(nums))
		copy(result, track)
		*results = append(*results, result)
	}

	for _, num := range nums {
		if !existIn(num, track) { //find options
			track = append(track, num)
			backtrack(nums, track, results) //go to next step
			track = track[0:len(track)-1] //remove options for this round (to choose a different option in next round)
		}
	}
}

func existIn(num int, track []int) bool {
	for _, x := range track {
		if num == x {
			return true
		}
	}
	return false
}