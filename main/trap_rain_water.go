package main

import "fmt"

/**
https://leetcode.com/problems/trapping-rain-water/

Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it can trap after raining.

Example 1:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9

Constraints:

n == height.length
0 <= n <= 3 * 104
0 <= height[i] <= 105
*/

func main() {
	fmt.Println("initializing test data...")
	fmt.Println("Answer is:", trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})) //6
	fmt.Println("Answer is:", trap([]int{4,2,0,3,2,5})) //9
}

/**
https://github.com/labuladong/fucking-algorithm/blob/master/%E9%AB%98%E9%A2%91%E9%9D%A2%E8%AF%95%E7%B3%BB%E5%88%97/%E6%8E%A5%E9%9B%A8%E6%B0%B4.md
 */
func trap(height []int) int {
	length := len(height)
	//can't possibly trap water
	if length < 3 {
		return 0
	}

	result := 0

	leftMax := height[0]
	rightMax := height[length-1]
	leftPointer := 0
	rightPointer := length - 1

	for leftPointer < rightPointer {
		if leftMax < rightMax {
			leftPointer ++
			if leftMax < height[leftPointer] {
				leftMax = height[leftPointer]
			}
			result += leftMax - height[leftPointer]
		} else {
			rightPointer --
			if rightMax < height[rightPointer] {
				rightMax = height[rightPointer]
			}
			result += rightMax - height[rightPointer]
		}
	}
	return result
}