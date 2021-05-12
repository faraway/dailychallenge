package main

import "fmt"

/**
https://leetcode.com/problems/permutations-ii/

Given a collection of numbers, nums, that might contain duplicates,
return all possible unique permutations in any order.



Example 1:
Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]

Example 2:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Constraints:
1 <= nums.length <= 8
-10 <= nums[i] <= 10

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", permuteII([]int{1,2,3})) //[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	fmt.Println("Answer is:", permuteII([]int{1,1,2})) //[[1,1,2],[1,2,1],[2,1,1]]
	fmt.Println("Answer is:", permuteII([]int{1}))     //[[1]]
}

var TOTAL_MAP map[int]int
var RESULTS [][]int

func permuteII(nums []int) [][]int {
	TOTAL_MAP = make(map[int]int)
	RESULTS = make([][]int,0,10)

	for _, num := range nums {
		TOTAL_MAP[num] += 1
	}
	backtrackPermuteII(nums, make([]int,0,len(nums)), make(map[int]int))
	return RESULTS
}

func backtrackPermuteII(nums []int, track []int, trackCounter map[int]int) {
	//add result to collection
	if len(track) == len(nums) {
		result := make([]int, len(nums))
		copy(result, track)
		RESULTS = append(RESULTS, result)
	}

	//in this for loop, we don't want to pick duplicate numbers to backtrack, hence the duplicate checker
	//e.g. if we first pick "1" during the loop, then next steps in this same loop I don't need to pick "1" again.
	//because they are the same number, and will result in same (duplicate) backtrack
	dupChecker := make(map[int]int)
	for _, num := range nums {
		if valid(num, trackCounter) && dupChecker[num] == 0 { //find options
			dupChecker[num] = 1
			track = append(track, num)
			trackCounter[num]++
			backtrackPermuteII(nums, track, trackCounter) //go to next step
			trackCounter[num]--
			track = track[0:len(track)-1] //remove options for this round (to choose a different option in next round)
		}
	}
}

func valid(num int, trackCounter map[int]int) bool {
	return trackCounter[num]+1 <= TOTAL_MAP[num]
}