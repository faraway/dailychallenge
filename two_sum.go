package main

import "fmt"

/**
You are given a list of numbers, and a target number k.
Return whether or not there are two numbers in the list that add up to k.

Try to do it in a single pass of the list.

Example:
Given [4, 7, 1 , -3, 2] and k = 5,
return true since 4 + 1 = 5.

*/

func main() {
	fmt.Println("initializing test data...")
	//TODO: Test cases
	fmt.Println("Answer is:", twoSum([]int{4,1,7.-3,2}, 5)) //True
	fmt.Println("Answer is:", twoSum([]int{4,1,7.-3,2}, 8)) //True
	fmt.Println("Answer is:", twoSum([]int{4,1,7.-3,2}, 15)) //False
	fmt.Println("Answer is:", twoSum([]int{0,1,7.-3,2}, 0)) //False
	fmt.Println("Answer is:", twoSum([]int{0,1,7.-3,0}, 0)) //True
}

func twoSum(array []int, target int) bool {
	checkMap := make(map[int]bool)
	for _, num := range array {
		if checkMap[target-num] {
			return true
		} else {
			checkMap[num] = true
		}
	}
	return false
}
