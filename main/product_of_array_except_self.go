package main

import "fmt"

/**
https://leetcode.com/problems/product-of-array-except-self/
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up:

Could you solve it in O(n) time complexity and without using division?
Could you solve it with O(1) constant space complexity? (The output array does not count as extra space for space complexity analysis.)

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", productOfArrayExceptSelf([]int{1,2,3,4})) //[24,12,8,6]
	fmt.Println("Answer is:", productOfArrayExceptSelf([]int{-1,1,0,-3,3})) //[0,0,9,0,0]
	fmt.Println("Answer is:", productOfArrayExceptSelf([]int{4,2})) //[2,4]

	fmt.Println("Answer is:", productOfArrayExceptSelfConstantSpace([]int{1,2,3,4})) //[24,12,8,6]
	fmt.Println("Answer is:", productOfArrayExceptSelfConstantSpace([]int{-1,1,0,-3,3})) //[0,0,9,0,0]
	fmt.Println("Answer is:", productOfArrayExceptSelfConstantSpace([]int{4,2})) //[2,4]
}

func productOfArrayExceptSelf(input []int) []int {
	size := len(input)
	//we don't need to do anything with [] or [x]; return the same
	if size < 2 {
		panic("invalid input. len(input) must be >= 2")
	}
	result := make([]int, size)

	//construct the supporting arrays 'left' and 'right'
	left := make([]int, size)
	right := make([]int, size)

	left[0] = input[0]
	right[size-1] = input[size-1]

	//left[i] = input[0] * input[1] * input[2] *....* input[i]
	//right[i] = input[size-1] * input[size-2] * input[size-3] *...* input[i]
	for i:=1; i<size; i++ {
		left[i] = left[i-1] * input[i]
		right[size-1-i] = right[size-i] * input[size-1-i]
	}

	//calculate result array
	result[0] = right[1]
	result[size-1] = left[size-2]
	for i:=1; i<size-1; i++ {
		result[i] = left[i-1] * right[i+1]
	}
	return result
}

/**
 This is an extension of the above idea, just try to save up space
 */
func productOfArrayExceptSelfConstantSpace(input []int) []int {
	size := len(input)
	//we don't need to do anything with [] or [x]; return the same
	if size < 2 {
		panic("invalid input. len(input) must be >= 2")
	}

	//similar array of 'left' above
	answer := make([]int, size)
	answer[0] = 1
	//answer[i] = input[0] * input[1] * input[2] *....* input[i-1]
	for i:=1; i<size; i++ {
		answer[i] = answer[i-1] * input[i-1]
	}

	//now calculate the 'right' part *on the fly*
	right := 1
	for i:=size-2; i>=0; i-- { //answer[size-1] is already taken care of above
		right *= input[i+1]
		answer[i] = answer[i] * right
	}
	return answer
}

