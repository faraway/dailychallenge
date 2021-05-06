package main

import (
	"fmt"
)

/**

https://leetcode.com/problems/set-mismatch/
You have a set of integers s, which originally contains all the numbers from 1 to n.
Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set,
which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

Example 1:
Input: nums = [1,2,2,4]
Output: [2,3]

Example 2:
Input: nums = [1,1]
Output: [1,2]

Constraints:
2 <= nums.length <= 104
1 <= nums[i] <= 104

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findErrorNums([]int{1,2,2,4})) //[2,3]
	fmt.Println("Answer is:", findErrorNums([]int{1,1})) //[1,2]
}

/**
O(n) time / O(1) space solution
ref: https://github.com/faraway/fucking-algorithm/blob/master/%E9%AB%98%E9%A2%91%E9%9D%A2%E8%AF%95%E7%B3%BB%E5%88%97/%E7%BC%BA%E5%A4%B1%E5%92%8C%E9%87%8D%E5%A4%8D%E7%9A%84%E5%85%83%E7%B4%A0.md
 */
func findErrorNums(nums []int) []int {
	// number is 1...n; array index is 0...n-1
	// use number-1 as index to mark the array number as negative;
	// this way we will see which number has been marked already and find duplicates
	dup := -1
	missing := -1
	for _, num := range nums {
		index := abs(num)-1
		if nums[index] < 0 {
			dup = index + 1
		} else {
			nums[index] = -1*nums[index]
		}
	}
	for i, num := range nums {
		if num > 0 {
			missing = i+1
			break
		}
	}
	return []int{dup,missing}
}

/**
 Go doesn't have math.abs for ints :(
 */
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
